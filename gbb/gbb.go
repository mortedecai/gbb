package gbb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/mortedecai/go-burn-bits/gbberror"
)

// GoBurnBits is the application interface to split the bulk of processing code out of the main package.
type GoBurnBits interface {
	Run([]string) error
	HandleDownload([]string) error
}

type GBBClient interface {
	Do(*http.Request) (*http.Response, error)
}

//go:generate mockgen -destination=./mocks/mock_client.go -package=mocks github.com/mortedecai/go-burn-bits/gbb GBBClient

type GBB struct {
	Host       string
	AuthToken  string
	Client     GBBClient
	WorkingDir string
}

const (
	CMD_DOWNLOAD = "download"
)

// New creates a GoBurnBits instance
func New(host string, token string) *GBB {
	wd, _ := os.Getwd()

	return &GBB{Host: host, AuthToken: token, Client: http.DefaultClient, WorkingDir: wd}
}

// Run starts the process of running the command line input
func (g *GBB) Run(args []string) error {
	if len(args) < 1 {
		return gbberror.ErrBadArguments
	}
	cmd := args[0]
	switch cmd {
	case CMD_DOWNLOAD:
		return g.HandleDownload(args[1:])
	default:
		return gbberror.ErrNotYetImplemented
	}
}

type GBBDownloadFilesResponse struct {
	Success bool                `json:"success"`
	Data    GBBFileDownloadData `json:"data,omitempty"`
}

type GBBFileDownloadData struct {
	Files []GBBDownloadFile `json:"files,omitempty"`
}

type GBBFileName string

func (gfn GBBFileName) IsValid() bool {
	return gfn != ""
}

func (gfn GBBFileName) HasDir() bool {
	return strings.Contains(gfn.String(), "/")
}

func (gfn GBBFileName) String() string {
	return string(gfn)
}

func (gfn GBBFileName) ToAbsolutePath(outputDir string) string {
	return path.Join(outputDir, path.Clean(gfn.String()))
}

func (gfn GBBFileName) CreatePath(outputDir string) error {
	absPath := gfn.ToAbsolutePath(outputDir)
	dirList := path.Dir(absPath)
	return os.MkdirAll(dirList, 0750)
}

type GBBDownloadFile struct {
	Filename GBBFileName `json:"filename"`
	Code     string      `json:"code,omitempty"`
	RamUsage int         `json:"ramUsage,omitempty"`
}

func (g *GBB) handleServerCall(req *http.Request, expStatus int, responseData any) error {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.AuthToken))

	resp, err := g.Client.Do(req)
	if err != nil {
		return fmt.Errorf(gbberror.StandardWrapper, gbberror.ErrRequestFailed, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != expStatus {
		return fmt.Errorf("%w: expected %d, got %d", gbberror.ErrUnexpectedResponse, http.StatusOK, resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(gbberror.StandardWrapper, gbberror.ErrResponseReadFailed, err.Error())
	}
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return fmt.Errorf(gbberror.StandardWrapper, gbberror.ErrBadJSON, err.Error())
	}

	return nil
}

// HandleDownload is responsible for parsing the necessary download arguments and fetching the files from the BitBurner server.
// If there is an issue with any of the arguments or the download an error will be returned. Nil on success.
func (g *GBB) HandleDownload(args []string) error {
	fmt.Printf("Starting download...")
	var outputDir string
	const (
		argAuthToken = "--authToken" //#nosec G101 -- This is a false positive
		argOutputDir = "--outputDir"
	)

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case argAuthToken:
			i++
			g.AuthToken = strings.TrimSpace(args[i])
		case argOutputDir:
			i++
			outputDir = strings.TrimSpace(args[i])
		}
	}
	if g.AuthToken == "" {
		return gbberror.ErrNoAuthToken
	}
	if len(outputDir) == 0 {
		return gbberror.ErrNoOutputDir
	}

	outputPath := path.Join(g.WorkingDir, path.Clean(outputDir))
	if !strings.HasPrefix(outputPath, g.WorkingDir) {
		return fmt.Errorf("%w: %s is outside working directory", gbberror.ErrBadOutputDir, outputPath)
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s", g.Host), bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return fmt.Errorf(gbberror.StandardWrapper, gbberror.ErrRequestFailed, err)
	}
	var downloadResults GBBDownloadFilesResponse
	if err = g.handleServerCall(req, http.StatusOK, &downloadResults); err != nil {
		return err
	}
	if !downloadResults.Success {
		return fmt.Errorf("%w: results file has success == false", gbberror.ErrBitBurnerFailure)
	}

	return g.WriteFiles(outputPath, downloadResults.Data.Files)
}

func (g *GBB) WriteFiles(outputDir string, files []GBBDownloadFile) error {
	failedFiles := make([]string, 0)

	for _, v := range files {
		if !v.Filename.IsValid() {
			continue
		}
		fp := v.Filename.ToAbsolutePath(outputDir)

		if v.Filename.HasDir() {
			if err := v.Filename.CreatePath(outputDir); err != nil {
				failedFiles = append(failedFiles, fmt.Sprintf("%s (%s)", v.Filename, fp))
				continue
			}
		}
		f, err := os.Create(path.Clean(fp))
		if err != nil {
			failedFiles = append(failedFiles, fmt.Sprintf("%s (%s)", v.Filename, fp))
			continue
		}
		err = g.writeFile(f, v)
		if err != nil {
			failedFiles = append(failedFiles, fmt.Sprintf("%s (%s)", v.Filename, fp))
			continue
		}
	}

	if len(failedFiles) > 0 {
		fmt.Printf("Failed to write %d files:\n", len(failedFiles))
		for i, v := range failedFiles {
			fmt.Printf("%d) %s\n", (i + 1), v)
		}
		return fmt.Errorf("%w: failed to write files", gbberror.ErrFileIssue)
	}

	return nil
}

func (g *GBB) writeFile(f *os.File, v GBBDownloadFile) error {
	defer f.Close()
	totalWritten := 0
	for attempts := 0; totalWritten < len(v.Code) && attempts < 10; attempts++ {
		amtWritten, err := f.WriteString(v.Code)
		if err != nil {
			return fmt.Errorf("%w: unable to write to file %s: %s", gbberror.ErrFileIssue, v.Filename, err.Error())
		}
		totalWritten += amtWritten
	}
	return nil
}
