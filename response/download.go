package response

import (
	"github.com/mortedecai/gbb/models"
)

type GBBDownloadFilesResponse struct {
	Success bool                `json:"success"`
	Data    GBBFileDownloadData `json:"data,omitempty"`
}

type GBBFileDownloadData struct {
	Files []models.GBBFileData `json:"files,omitempty"`
}
