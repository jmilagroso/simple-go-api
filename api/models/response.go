package models

// ErrorResponse error response
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
