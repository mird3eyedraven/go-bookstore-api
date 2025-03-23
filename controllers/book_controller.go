package controllers

import (
	"go-bookstore-api/config"
	"go-bookstore-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}
