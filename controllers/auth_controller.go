package controllers

import (
	"go-bookstore-api/config"
	"go-bookstore-api/models"
	"go-bookstore-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username might already exist"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}
