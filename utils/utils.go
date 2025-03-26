package utils

import (
	"expense-tracker-with-go/models"
	"fmt"

	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Makan"},
		{Name: "Transportasi"},
		{Name: "Gaji"},
		{Name: "Hiburan"},
	}

	for _, category := range categories {
		var existingCategory models.Category
		if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			db.Create(&category)
			fmt.Println("Seeded category:", category.Name)
		} else {
			fmt.Println("Category already exists:", existingCategory.Name)
		}
	}
}