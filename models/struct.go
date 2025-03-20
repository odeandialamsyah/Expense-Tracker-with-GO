package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string `json:"username" gorm:"unique;not null"`
	Email		string `json:"email" gorm:"unique;not null"`
	Password 	string `json:"password" gorm:"not null"`
	RoleID   uint   `json:"role_id" gorm:"not null"`
    Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}

type Role struct {
    gorm.Model
    Name  string `json:"name" gorm:"unique;not null"`
    Users []User `gorm:"foreignKey:RoleID" json:"-"`
}

type Transaction struct {
    gorm.Model
    Amount      float64 `json:"amount" gorm:"not null"`
    Description string  `json:"description" gorm:"not null"`
    Date        string  `json:"date" gorm:"not null"`
    CategoryID  uint    `json:"category_id" gorm:"not null"`
    Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}

type Category struct {
    gorm.Model
    Name         string       `json:"name" gorm:"unique;not null"`
    Transactions []Transaction `json:"transactions" gorm:"foreignKey:CategoryID"`
}

func MigrateUsers(db *gorm.DB) {
    db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&Transaction{})
    db.AutoMigrate(&Category{})
}