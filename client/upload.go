package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mortedecai/gbb/gbberror"
	"github.com/mortedecai/gbb/models"
	"github.com/mortedecai/gbb/response"
	"net/http"
	"os"
	"path"
)

func HandleUpload(uo UploadOption) error {
	var req *http.Request
	var buff *bytes.Reader
	var err error

	if buff, err = assembleUploadData(uo); err != nil {
		return err
	}

	if req, err = http.NewRequest(http.MethodPut, uo.Server(), buff); err != nil {
		return err
	}
	req = uo.AddAuth(req)

	var responseData response.GBBUploadFileResponse
	if err = handleServerCall(req, http.StatusOK, &responseData); err != nil {
		return fmt.Errorf("%w: %s", gbberror.ErrRequestFailed, err.Error())
	}

	if !responseData.Success {
		return fmt.Errorf("%w: success response was false", gbberror.ErrBitBurnerFailure)
	}
	return nil
}

func assembleUploadData(uo UploadOption) (*bytes.Reader, error) {
	var err error
	var fd models.GBBFileData
	var data []byte
	fd.Filename = uo.ToUpload()[0]

	if fd.Code, err = readFileData(fd.Filename); err != nil {
		return nil, err
	}
	fd.Code = base64.RawStdEncoding.EncodeToString([]byte(fd.Code))
	data, err = json.Marshal(fd)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func readFileData(fn models.GBBFileName) (string, error) {
	b, err := os.ReadFile(path.Clean(fn.String()))
	return string(b), err
}
