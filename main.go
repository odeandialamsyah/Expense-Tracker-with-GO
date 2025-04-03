package main

import (
	"expense-tracker-with-go/config"
	"expense-tracker-with-go/models"
	"expense-tracker-with-go/routes"
	"expense-tracker-with-go/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	models.Migrate(config.DB)
	utils.SeedCategories(config.DB)

	r := gin.Default()
	routes.TransactionRoutes(r, config.DB)

	r.Run(":8080")
}