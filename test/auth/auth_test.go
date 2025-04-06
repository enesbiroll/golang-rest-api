package services

import (
	"rest-api/models"
	"rest-api/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	// Test case 1: Register a new user successfully
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Assuming DB is already mocked or initialized properly in config.DB
	err := services.RegisterUser(user)
	assert.Nil(t, err, "Expected no error, but got %v", err)

	// Test case 2: Trying to register an existing user
	existingUser := &models.User{
		Username: "existinguser",
		Email:    "test@example.com", // Same email as before
		Password: "password123",
	}
	err = services.RegisterUser(existingUser)
	assert.NotNil(t, err, "Expected error due to existing email, but got nil")
	assert.Equal(t, "user with this email already exists", err.Error(), "Unexpected error message")
}

func TestLoginUser(t *testing.T) {
	// First, create a user to test login
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Assuming RegisterUser is called before testing login
	_ = services.RegisterUser(user)

	// Test case 1: Valid login
	loginUser, err := services.LoginUser("test@example.com", "password123") // Corrected to services.LoginUser
	assert.Nil(t, err, "Expected no error, but got %v", err)
	assert.NotNil(t, loginUser, "Expected user to be returned")

	// Test case 2: Invalid email
	_, err = services.LoginUser("invalid@example.com", "password123") // Corrected to services.LoginUser
	assert.NotNil(t, err, "Expected error due to invalid email, but got nil")
	assert.Equal(t, "user not found", err.Error(), "Unexpected error message")

	// Test case 3: Invalid password
	_, err = services.LoginUser("test@example.com", "wrongpassword") // Corrected to services.LoginUser
	assert.NotNil(t, err, "Expected error due to incorrect password, but got nil")
	assert.Equal(t, "invalid credentials", err.Error(), "Unexpected error message")
}

func TestDeleteUser(t *testing.T) {
	// First, create a user to test deletion
	user := &models.User{
		Username: "deleteuser",
		Email:    "delete@example.com",
		Password: "password123",
	}

	// Assuming RegisterUser is called before testing delete
	_ = services.RegisterUser(user)

	// Test case 1: Delete user successfully
	err := services.DeleteUser(user.Email) // Corrected to services.DeleteUser
	assert.Nil(t, err, "Expected no error, but got %v", err)

	// Test case 2: Try deleting a user that doesn't exist
	err = services.DeleteUser("nonexistent@example.com")
	assert.NotNil(t, err, "Expected error for deleting non-existent user, but got nil")
	assert.Equal(t, "user not found", err.Error(), "Unexpected error message")
}
