package repository

import (
	"expense-tracker-with-go/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	return ur.DB.Save(user).Error
}

func (ur *UserRepository) DeleteUser(id uint) error {
	return ur.DB.Delete(&models.User{}, id).Error
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}