package repository

import (
	"expense-tracker-with-go/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetAllTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Find(&transactions).Error
	return transactions, err
}