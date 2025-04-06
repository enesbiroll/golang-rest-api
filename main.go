package main

import (
	"rest-api/config"
	"rest-api/core/logger"
	_ "rest-api/docs" // Import the docs package for Swagger
	"rest-api/routes"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger" // Import Swagger middleware
)

func main() {
	// Initialize logger
	logger.Init()

	// Connect to database
	config.Connect()

	// Create a new Fiber app
	app := fiber.New()

	// Set up routes
	routes.StudentRoute(app)
	routes.AuthRoute(app)

	// Serve Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler) // Serve Swagger UI at /swagger/*

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		logger.Log.Fatalf("Error starting server: %v", err)
	}
}
