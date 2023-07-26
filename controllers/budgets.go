package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-accounting/models"
)

type CreateBudgetParam struct {
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
}

type UpdateBudgetParam struct {
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	TotalBudgeted     float64   `json:"totalBudgeted"`
	TotalAllocated    float64   `json:"totalAllocated"`
	BalanceToAllocate float64   `json:"balanceToAllocate"`
}

// GET /budgets
// Get all budgets
func FindBudgets(c *gin.Context) {
	var budgets []models.Budget
	models.DB.Find(&budgets)

	c.JSON(http.StatusOK, gin.H{"data": budgets})
}

// POST /budgets
// Create a new budget
func CreateBudget(c *gin.Context) {
	// Validate input
	var input CreateBudgetParam
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create budget
	budget := models.Budget{StartDate: input.StartDate, EndDate: input.EndDate}
	models.DB.Create(&budget)

	c.JSON(http.StatusOK, gin.H{"data": budget})
}

// GET /budgets/:id
// Find a budget
func FindBudget(c *gin.Context) { // Get model if exist
	var budget models.Budget

	if err := models.DB.Where("id = ?", c.Param("id")).First(&budget).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budget})
}

// PATCH /budgets/:id
// Update a budget
func UpdateBudget(c *gin.Context) {
	// Get model if exist
	var budget models.Budget
	if err := models.DB.Where("id = ?", c.Param("id")).First(&budget).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBudgetParam
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&budget).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": budget})
}

// DELETE /budgets/:id
// Delete a budget
func DeleteBudget(c *gin.Context) {
	// Get model if exist
	var budget models.Budget
	if err := models.DB.Where("id = ?", c.Param("id")).First(&budget).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&budget)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
