package routes

import (
	"rest-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	// Register Route
	app.Post("/auth/register", controllers.Register)

	// Login Route
	app.Post("/auth/login", controllers.Login)
}
