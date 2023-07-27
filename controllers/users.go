package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-accounting/models"
)

type CreateUserParam struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type UpdateUserParam struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// POST /signup
// Create a new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserParam
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := HashPassword(input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if match := CheckPasswordHash(input.ConfirmPassword, hash); !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Username: input.Username, Password: hash}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GET /current_user
// Get current user
func GetCurrentUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("username = ?", CurrentUser(c)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
		return
	}

	var userWithoutPassword = models.UserWithoutPassword{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	c.JSON(http.StatusOK, gin.H{"data": userWithoutPassword})
}

// PATCH /current_user
// Update existing user
func UpdateUser(c *gin.Context) {
	// Validate input
	var input UpdateUserParam
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("username = ?", CurrentUser(c)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
		return
	}

	models.DB.Model(&user).Updates(input)

	var userWithoutPassword = models.UserWithoutPassword{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	c.JSON(http.StatusOK, gin.H{"data": userWithoutPassword})
}
