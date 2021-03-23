package api

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RespondWithSuccess
func RespondWithSuccess(data interface{}, message string) SuccessResponse {
	response := SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	return response
}

// ReturnResponse
func ReturnResponse(ctx *gin.Context, response interface{}, statusCode int) {
	ctx.JSON(statusCode, response)
}

