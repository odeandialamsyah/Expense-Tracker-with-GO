package controllers

import (
	"expense-tracker-with-go/models"
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

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
    var transaction models.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi kategori
    _, err := tc.CategoryRepo.GetCategoryByID(transaction.CategoryID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    if err := tc.TransactionRepo.CreateTransaction(&transaction); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
        return
    }
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully", "data": transaction})
}
