package repository

import (
	"expense-tracker-with-go/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (rr *RoleRepository) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	if err := rr.DB.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}