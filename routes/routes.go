package routes

import (
	"go-bookstore-api/controllers"
	"go-bookstore-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)

	book := r.Group("/books")
	book.Use(middleware.JWTAuth())
	book.GET("/", controllers.GetBooks)
	book.POST("/", controllers.CreateBook)
	book.PUT("/:id", controllers.UpdateBook)
	book.DELETE("/:id", controllers.DeleteBook)
}
