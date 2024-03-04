package routes

import (
	"golang-sheets/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	r.GET("/download", handlers.DownloadExpensesAsCSV(db))
	r.GET("/expenses", handlers.ListExpenses(db))
	r.POST("/expenses", handlers.CreateExpense(db))
	r.GET("/expenses/:id", handlers.GetExpense(db))
	r.PUT("/expenses/:id", handlers.UpdateExpense(db))
	r.DELETE("/expenses/:id", handlers.DeleteExpense(db))

}
