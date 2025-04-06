package middleware

import (
	jwt "rest-api/core"
	"rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware korumalı route'lar için JWT doğrulaması
func AuthMiddleware(c *fiber.Ctx) error {
	// Header'dan token al
	token := c.Get("Authorization")
	if token == "" {
		return utils.ErrorResponse(c, "No token provided")
	}

	// Token'ı doğrula
	user, err := jwt.ParseJWT(token)
	if err != nil {
		return utils.ErrorResponse(c, "Invalid token")
	}

	// Kullanıcı bilgisini context'e ekle
	c.Locals("user", user)

	return c.Next()
}
