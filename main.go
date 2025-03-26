package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"inventory-management-api/db"
	"inventory-management-api/routes"
)

func main() {
	log.Println("Starting server on :8080")
	db.Init()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run()
}