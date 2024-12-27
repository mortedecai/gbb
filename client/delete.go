package client

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/mortedecai/gbb/gbberror"
	"github.com/mortedecai/gbb/response"
)

// HandleDelete is responsible for parsing the necessary download arguments and fetching the files from the BitBurner server.
// If there is an issue with any of the arguments or the download an error will be returned. Nil on success.
func HandleDelete(do DeleteOption) error {
	buff := bytes.NewReader([]byte(fmt.Sprintf(`{"filename":"%s"}`, do.ToDelete())))

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s:%d", do.Host(), do.Port()), buff)
	if err != nil {
		return fmt.Errorf(gbberror.StandardWrapper, gbberror.ErrRequestFailed, err)
	}
	req = do.AddAuth(req)

	var deleteResult response.GBBDeleteFileResponse
	if err = handleServerCall(req, http.StatusOK, &deleteResult); err != nil {
		return err
	}
	if !deleteResult.Success {
		return fmt.Errorf("%w: results file has success == false; message: '%s'", gbberror.ErrBitBurnerFailure, deleteResult.Message)
	}

	return nil
}
