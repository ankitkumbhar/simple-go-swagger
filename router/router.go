package router

import (
	"github.com/gin-gonic/gin"
	"go-swagger/controllers"
)

// PrepareRoutes
func PrepareRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.GET("/posts", controllers.Index)
}
