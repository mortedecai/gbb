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
	Host      string
	AuthToken string
	Client    GBBClient
}

const (
	CMD_DOWNLOAD = "download"
)

// New creates a GoBurnBits instance
func New(host string, token string) *GBB {
	return &GBB{Host: host, AuthToken: token, Client: http.DefaultClient}
}

// Run starts the process of running the command line input
func (g *GBB) Run(_ []string) error {
	return gbberror.NotYetImplemented
}

type GBBDownloadFilesResponse struct {
	Success bool                `json:"success"`
	Data    GBBFileDownloadData `json:"data,omitempty"`
}

type GBBFileDownloadData struct {
	Files []GBBDownloadFile `json:"files,omitempty"`
}

type GBBDownloadFile struct {
	Filename string `json:"filename"`
	Code     string `json:"code,omitempty"`
	RamUsage int    `json:"ramUsage,omitempty"`
}

// HandleDownload is responsible for parsing the necessary download arguments and fetching the files from the BitBurner server.
// If there is an issue with any of the arguments or the download an error will be returned. Nil on success.
func (g *GBB) HandleDownload(args []string) error {
	var authToken string
	var outputDir string
	const (
		argAuthToken = "--authToken"
		argOutputDir = "--outputDir"
	)

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case argAuthToken:
			i++
			authToken = args[i]
		case argOutputDir:
			i++
			outputDir = strings.TrimSpace(args[i])
		}
	}
	if strings.TrimSpace(authToken) == "" {
		return gbberror.NoAuthToken
	}
	if len(outputDir) == 0 {
		return gbberror.NoOutputDir
	}
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%w: working dir %s", gbberror.FileIssue, err.Error())
	}

	outputPath := path.Join(wd, outputDir)
	if !strings.HasPrefix(outputPath, wd) {
		return fmt.Errorf("%w: %s is outside working directory", gbberror.BadOutputDir, outputPath)
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s", g.Host), bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := g.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: expected %d, got %d", gbberror.UnexpectedResponse, http.StatusOK, resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%w: %s", gbberror.ResponseReadFailed, err.Error())
	}
	var downloadResults GBBDownloadFilesResponse
	err = json.Unmarshal(data, &downloadResults)
	if err != nil {
		return fmt.Errorf("%w: %s", gbberror.BadJSON, err.Error())
	}

	if !downloadResults.Success {
		return fmt.Errorf("%w: results file has success == false", gbberror.BitBurnerFailure)
	}

	failedFiles := make([]string, 0)

	for _, v := range downloadResults.Data.Files {
		fp := path.Join(outputPath, v.Filename)
		f, err := os.Create(fp)
		if err != nil {
			failedFiles = append(failedFiles, fmt.Sprintf("%s (%s)", v.Filename, fp))
			continue
		}
		g.writeFile(f, v)
	}

	if len(failedFiles) > 0 {
		fmt.Printf("Failed to write %d files:\n", len(failedFiles))
		for i, v := range failedFiles {
			fmt.Printf("%d) %s\n", (i + 1), v)
		}
		return fmt.Errorf("%w: failed to write files", gbberror.FileIssue)
	}

	return nil
}

func (g *GBB) writeFile(f *os.File, v GBBDownloadFile) error {
	defer f.Close()
	totalWritten := 0
	for attempts := 0; totalWritten < len(v.Code) && attempts < 10; attempts++ {
		amtWritten, err := f.WriteString(v.Code)
		if err != nil {
			return fmt.Errorf("%w: unable to write to file %s: %s", gbberror.FileIssue, v.Filename, err.Error())
		}
		totalWritten += amtWritten
	}
	return nil
}
