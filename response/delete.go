package response

// GBBDeleteFileResponse is used to capture the BitBurner server response on a file deletion attempt.
type GBBDeleteFileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"msg,omitempty"`
}
