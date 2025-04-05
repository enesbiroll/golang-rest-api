package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Handle panic recovery
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Something went wrong, please try again later",
				})
			}
		}()

		// Continue processing the request
		err := c.Next()
		if err != nil {
			log.Printf("Error occurred: %v", err)
			// Handle specific errors (you can extend this)
			if e, ok := err.(*fiber.Error); ok {
				// Specific custom error
				return c.Status(e.Code).JSON(fiber.Map{
					"status":  "error",
					"message": e.Message,
				})
			}
			// Default internal server error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Internal server error",
			})
		}
		return nil
	}
}
