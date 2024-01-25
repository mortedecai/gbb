package gbb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mortedecai/gbb/gbberror"
	"github.com/mortedecai/gbb/response"
	"io"
	"net/http"
	"os"
	"path"
)

// GBBClient is an interface to the http.Client methods used for mocking purposes
type GBBClient interface {
	Do(*http.Request) (*http.Response, error)
}

//go:generate mockgen -destination=./mocks/mock_client.go -package=mocks github.com/mortedecai/gbb/gbb GBBClient

var (
	Client GBBClient = http.DefaultClient
)

func handleServerCall(req *http.Request, expStatus int, responseData any) error {
	resp, err := Client.Do(req)
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
func HandleDownload(do DownloadOption) error {
	fmt.Printf("Starting download...")
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%d", do.Host(), do.Port()), bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return fmt.Errorf(gbberror.StandardWrapper, gbberror.ErrRequestFailed, err)
	}
	req = do.AddAuth(req)

	var downloadResults response.GBBDownloadFilesResponse
	if err = handleServerCall(req, http.StatusOK, &downloadResults); err != nil {
		return err
	}
	if !downloadResults.Success {
		return fmt.Errorf("%w: results file has success == false", gbberror.ErrBitBurnerFailure)
	}

	return WriteFiles(path.Clean(do.Destination()), downloadResults.Data.Files)
}

func WriteFiles(outputDir string, files []response.GBBDownloadFile) error {
	failedFiles := make([]string, 0)
	const failedFileStr = "%s (%s)"

	for _, v := range files {
		if !v.Filename.IsValid() {
			continue
		}
		fp := v.Filename.ToAbsolutePath(outputDir)

		if v.Filename.HasDir() {
			if err := v.Filename.CreatePath(outputDir); err != nil {
				failedFiles = append(failedFiles, fmt.Sprintf(failedFileStr, v.Filename, fp))
				continue
			}
		}
		f, err := os.Create(path.Clean(fp))
		if err != nil {
			failedFiles = append(failedFiles, fmt.Sprintf(failedFileStr, v.Filename, fp))
			continue
		}
		err = writeFile(f, v)
		if err != nil {
			failedFiles = append(failedFiles, fmt.Sprintf(failedFileStr, v.Filename, fp))
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

func writeFile(f *os.File, v response.GBBDownloadFile) error {
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
