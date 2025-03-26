package routes

import (
	"inventory-management-api/controllers"
	"inventory-management-api/middleware"
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

	// resources for inventory
	inv := r.Group("/inventory")
	inv.GET("/", controllers.ListItems)

	// resources that require authenticaion
	inv.Use(middleware.RequireAuth)
	{
		inv.POST("/", controllers.CreateItem)
		inv.POST("/:id/restock", controllers.RestockItem)
		inv.GET("/:id/restocks", controllers.RestockHistory)
	}
}
