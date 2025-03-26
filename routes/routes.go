package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the root route!")
	})

	// Public endpoints for registration and login
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
