package controllers

import (
	jwt "rest-api/core"
	"rest-api/core/logger"
	"rest-api/models"
	"rest-api/services"
	"rest-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Register handles the registration of a new user
// @Summary Register a new user
// @Description Registers a new user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "User Information"
// @Success 200 {string} string "User registered successfully"
// @Failure 400 {string} string "Invalid input"
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		// Log the error for debugging
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to parse request body")

		return utils.ErrorResponse(c, "Invalid input")
	}

	// Validate that fields are not empty
	if user.Email == "" || user.Password == "" || user.Username == "" {
		return utils.ErrorResponse(c, "Missing required fields")
	}

	// Log the registration request for debugging
	logger.Log.WithFields(logrus.Fields{
		"email": user.Email,
	}).Info("Register request received")

	// Call the service to register the user
	if err := services.RegisterUser(&user); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"email": user.Email,
		}).Error("Registration failed")
		return utils.ErrorResponse(c, err.Error())
	}

	// Return a success response
	return utils.SuccessResponseNoData(c, "User registered successfully")
}

// Login handles user login and JWT token generation
// @Summary Login a user and get JWT token
// @Description Logs in with email and password, returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "User Login Information"
// @Success 200 {object} map[string]string "Login successful with token"
// @Failure 400 {string} string "Invalid credentials"
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		// Log the error for debugging
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to parse request body")
		return utils.ErrorResponse(c, "Invalid input")
	}

	// Log the login attempt for debugging
	logger.Log.WithFields(logrus.Fields{
		"email": user.Email,
	}).Info("Login request received")

	// Validate the user credentials by calling the service
	foundUser, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		// Log the failed login attempt
		logger.Log.WithFields(logrus.Fields{
			"email": user.Email,
		}).Error("Login failed: Invalid credentials")
		return utils.ErrorResponse(c, err.Error())
	}

	// Generate JWT for the user
	token, err := jwt.GenerateJWT(foundUser)
	if err != nil {
		return utils.ErrorResponse(c, "Error generating token")
	}

	// Return success response with the token
	return utils.SuccessResponse(c, "Login successful", map[string]string{
		"token": token,
	})
}
