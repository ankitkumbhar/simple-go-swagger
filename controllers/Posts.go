package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go-swagger/api"
)

type (
	Posts struct {
		Name        string    `json:"name" example:"A"`
		Description string    `json:"description" example:"A"`
		Category    string    `json:"category" example:"A"`
		CreatedAt   time.Time `json:"created_at" example:"2019-11-09T21:21:46+00:00"`
		UpdatedAt   time.Time `json:"updated_at" example:"2019-11-09T21:21:46+00:00"`
	}

	PostResponse struct {
		Success bool    `json:"success" example:"true"`
		Message string  `json:"message" example:""`
		Data    []Posts `json:"data"`
	}
)

// Index godoc
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} PostResponse
// @Failure 400 {object} interface{}
// @Failure 404 {object} interface{}
// @Failure 500 {object} interface{}
// @Router /posts [get]
func Index(ctx *gin.Context) {
	posts := []Posts{{
		Name:        "Test name",
		Description: "Test Description",
		Category:    "Test Category",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}}
	resp := api.RespondWithSuccess(posts, "")
	api.ReturnResponse(ctx, resp, http.StatusOK)
}
