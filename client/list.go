package client

import (
	"bytes"
	"fmt"
	"github.com/mortedecai/gbb/gbberror"
	"github.com/mortedecai/gbb/response"
	"net/http"
	"strings"
)

func HandleList(lo CommandOption) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%d", lo.Host(), lo.Port()), bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return fmt.Errorf(gbberror.StandardWrapper, gbberror.ErrRequestFailed, err)
	}
	req = lo.AddAuth(req)

	var downloadResults response.GBBDownloadFilesResponse
	if err = handleServerCall(req, http.StatusOK, &downloadResults); err != nil {
		return err
	}
	if !downloadResults.Success {
		return fmt.Errorf("%w: results file has success == false", gbberror.ErrBitBurnerFailure)
	}
	builder := strings.Builder{}
	builder.WriteString("BitBurner Files\n")
	builder.WriteString("===============\n")
	builder.WriteString("\n")
	for _, v := range downloadResults.Data.Files {
		builder.WriteString(v.Filename.String())
		builder.WriteString("\n")
	}
	fmt.Println(builder.String())
	return nil
}
