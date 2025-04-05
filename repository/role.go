package repository

import (
	"expense-tracker-with-go/models"

	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) roleRepository {
	return roleRepository{DB: db}
}

func (rr *roleRepository) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	if err := rr.DB.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}