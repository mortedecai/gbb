package client

import (
	"github.com/mortedecai/gbb/gbberror"
	"net/http"
)

func HandleUpload(uo UploadOption) error {
	_, err := http.NewRequest(http.MethodPut, uo.Server(), nil)
	if err != nil {
		return err
	}
	return gbberror.ErrNotYetImplemented
}
