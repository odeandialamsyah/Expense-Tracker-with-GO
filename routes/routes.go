package routes

import (
	"expense-tracker-with-go/controllers"
	"expense-tracker-with-go/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionRoutes(router *gin.Engine, db *gorm.DB) {
	//initialize repo
	transactionRepo := repository.NewTransactionRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	//initialize controller with repo
	tc := controllers.TransactionController {
		TransactionRepo: *transactionRepo,
		CategoryRepo: *categoryRepo,
	}

	//define routes
	transactionGroup := router.Group("/transaction") 
	{
		transactionGroup.GET("", tc.GetTransactions)
		transactionGroup.POST("", tc.CreateTransaction)
		transactionGroup.PUT("/:id", tc.UpdateTransaction)
		transactionGroup.DELETE("/:id", tc.DeleteTransaction)
	}
}