package repository

import (
	"expense-tracker-with-go/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{DB: db}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}