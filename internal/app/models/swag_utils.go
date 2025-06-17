package models

// ErrorResponse represents an error message response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DefaultResponse represents a default message response
type DefaultResponse struct {
	Message string `json:"message"`
}
