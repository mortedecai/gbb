package client

import (
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

//go:generate mockgen -destination=./mocks/mock_client.go -package=mocks github.com/mortedecai/gbb/client GBBClient

type fileWriterFunc func(path string) (FileWriter, error)

type FileWriter interface {
	WriteString(data string) (n int, err error)
	Close() error
}

//go:generate mockgen -destination=./mocks/file_writer.go -package=mocks github.com/mortedecai/gbb/client FileWriter

var (
	Client     GBBClient      = http.DefaultClient
	createFile fileWriterFunc = func(path string) (FileWriter, error) { return os.Create(path) } // #nosec G304
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
		//f, err := os.Create(path.Clean(fp))
		f, err := createFile(path.Clean(fp))
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

func writeFile(f FileWriter, v response.GBBDownloadFile) error {
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
