package client

import (
	"bytes"
	"fmt"
	"github.com/mortedecai/gbb/gbberror"
	"github.com/mortedecai/gbb/response"
	"net/http"
	"path"
)

// HandleDownload is responsible for parsing the necessary download arguments and fetching the files from the BitBurner server.
// If there is an issue with any of the arguments or the download an error will be returned. Nil on success.
func HandleDownload(do DownloadOption) error {
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
