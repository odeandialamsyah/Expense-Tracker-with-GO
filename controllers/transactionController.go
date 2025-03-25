package controllers

import (
	"expense-tracker-with-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionRepo repository.TransactionRepository
	CategoryRepo    repository.CategoryRepository
}

func (tc *TransactionController) GetTransactions(c *gin.Context) {
	transactions, err := tc.TransactionRepo.GetAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
