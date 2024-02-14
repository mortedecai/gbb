package response

type GBBFailureResponse struct {
	Success bool   `json:"success"`
	Message string `json:"msg,omitempty"`
}
