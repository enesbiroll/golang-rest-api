package main

import (
	Config "rest-api/config"
	"rest-api/core/logger"
	routes "rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the logger
	logger.Init()

	// Initialize the database connection
	Config.Connect()

	// Create a new Fiber app instance
	app := fiber.New()

	// Set up the routes for the app
	routes.StudentRoute(app)

	// Start the application on port 3000
	if err := app.Listen(":3000"); err != nil {
		logger.Log.Fatalf("Error starting server: %v", err)
	}
	// Log to console and file
	logger.Log.Info("This is an info message")
	logger.Log.Warn("This is a warning message")
	logger.Log.Fatal("This is a fatal error message")
}
