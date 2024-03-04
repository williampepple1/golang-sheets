package handlers

import (
	"golang-sheets/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListExpenses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var expenses []models.Expense
		if err := db.Find(&expenses).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve expenses"})
			return
		}
		c.JSON(http.StatusOK, expenses)
	}
}

func CreateExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var expense models.Expense
		if err := c.ShouldBindJSON(&expense); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&expense).Error; err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed to create expense"})
			return
		}

		c.JSON(http.StatusCreated, expense)
	}
}

func GetExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var expense models.Expense
		if err := db.Where("id = ?", id).First(&expense).Error; err != nil {
			c.JSON(404, gin.H{"error": "Expense not found"})
			return
		}

		c.JSON(http.StatusOK, expense)
	}
}

func UpdateExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var expense models.Expense
		if err := db.Where("id = ?", id).First(&expense).Error; err != nil {
			c.JSON(404, gin.H{"error": "Expense not found"})
			return
		}

		if err := c.ShouldBindJSON(&expense); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Save(&expense).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update expense"})
			return
		}

		c.JSON(200, expense)
	}
}

func DeleteExpense(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var expense models.Expense
		if err := db.Where("id = ?", id).First(&expense).Error; err != nil {
			c.JSON(404, gin.H{"error": "Expense not found"})
			return
		}

		if err := db.Delete(&expense).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete expense"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
	}
}
