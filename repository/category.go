package repository

import (
	"expense-tracker-with-go/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := cr.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr *CategoryRepository) CreateCategory(category *models.Category) error {
	return cr.db.Create(category).Error
}

func (cr *CategoryRepository) UpdateCategory(category *models.Category) error {
	return cr.db.Save(category).Error
}

func (cr *CategoryRepository) DeleteCategory(id uint) error {
	return cr.db.Delete(&models.Category{}, id).Error
}