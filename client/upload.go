package client

import (
	"bytes"
	"encoding/json"
	"github.com/mortedecai/gbb/gbberror"
	"github.com/mortedecai/gbb/models"
	"net/http"
	"os"
	"path"
)

func HandleUpload(uo UploadOption) error {
	buff, err := assembleUploadData(uo)
	if err != nil {
		return err
	}
	_, err = http.NewRequest(http.MethodPut, uo.Server(), buff)
	if err != nil {
		return err
	}
	return gbberror.ErrNotYetImplemented
}

func assembleUploadData(uo UploadOption) (*bytes.Reader, error) {
	var err error
	var fd models.GBBFileData
	var data []byte
	fd.Filename = uo.ToUpload()[0]

	if fd.Code, err = readFileData(fd.Filename); err != nil {
		return nil, err
	}
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
