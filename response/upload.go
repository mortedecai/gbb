package response

type GBBUploadFileResponse struct {
	Success bool              `json:"success"`
	Data    GBBFileUploadData `json:"data"`
}

type GBBFileUploadData struct {
	Overwritten bool `json:"overwritten"`
	RamUsage    *int `json:"ramUsage"`
}
