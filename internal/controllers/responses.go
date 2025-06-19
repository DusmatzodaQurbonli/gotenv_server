package controllers

import "Gotenv/internal/app/models"

func newErrorResponse(error string) models.ErrorResponse {
	return models.ErrorResponse{
		Error: error,
	}
}
