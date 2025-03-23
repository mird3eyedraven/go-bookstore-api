package main

import (
	"go-bookstore-api/config"
	"go-bookstore-api/mird3eyedraven/routes"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	routes.RegisterRoutes(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
