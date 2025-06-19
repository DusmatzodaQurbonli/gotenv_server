package models

// ErrorResponse represents an error message response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DefaultResponse represents a default message response
type DefaultResponse struct {
	Message string `json:"message"`
}

// TokenResponse represents the response with access token and user ID
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       uint   `json:"user_id"`
}

// RefreshTokenResponse represents the response with access token and user ID
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
