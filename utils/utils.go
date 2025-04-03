package utils

import (
	"bytes"
	"expense-tracker-with-go/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err 
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil 
}