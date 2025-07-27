package controllers

import (
	"net/http"

	"github.com/Truong62/taskoria/auth-service/config"
	"github.com/Truong62/taskoria/auth-service/models"
	"github.com/Truong62/taskoria/auth-service/utils"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Hash password error"})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: hashedPassword,
		FullName: input.FullName,
		Role:     "user",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Register successfully"})
}
