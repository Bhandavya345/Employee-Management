package main

import (
	"log"

	"github.com/Bhandavya345/Employee-Management/database"
	"github.com/Bhandavya345/Employee-Management/logger"
	"github.com/Bhandavya345/Employee-Management/models"
	"github.com/Bhandavya345/Employee-Management/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Logger
	logger.InitLogger()
	logger.InfoLogger.Println("Starting Employee Management Application...")

	// Connect Database
	database.ConnectDB()
	logger.InfoLogger.Println("Database Connected Successfully")

	// Auto Migrate Tables
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Employee{},
	)

	if err != nil {
		logger.ErrorLogger.Println("AutoMigrate Failed:", err)
		log.Fatal(err)
	}

	logger.InfoLogger.Println("Database Migration Completed")

	// Create Gin Router
	router := gin.Default()

	// Register Routes
	routes.SetupRoutes(router)

	logger.InfoLogger.Println("Routes Registered Successfully")

	// Start Server
	logger.InfoLogger.Println("Server Running On http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		logger.ErrorLogger.Println("Failed To Start Server:", err)
		log.Fatal(err)
	}
}
