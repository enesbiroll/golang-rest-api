package services

import (
	"errors"
	"rest-api/config"
	"rest-api/models"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registers a new user by hashing the password and saving to DB
func RegisterUser(user *models.User) error {
	// Check if the user already exists by email
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("user with this email already exists")
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}

	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := config.DB.Create(user).Error; err != nil {
		return errors.New("error creating user")
	}

	return nil
}

// LoginUser checks if the provided email and password match any user in DB
func LoginUser(email, password string) (*models.User, error) {
	// Find the user by email
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Compare the stored hashed password with the entered password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

// DeleteUser deletes a user by email (soft delete)
func DeleteUser(email string) error {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	// Soft delete the user (this assumes the DB supports soft deletes, for example by having a "deleted_at" field)
	if err := config.DB.Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
