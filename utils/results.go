package utils

import "github.com/gofiber/fiber/v2"

// Success response with data
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	result := NewSuccessDataResult(message, data)
	return c.Status(fiber.StatusOK).JSON(result)
}

// Success response without data
func SuccessResponseNoData(c *fiber.Ctx, message string) error {
	result := NewSuccessResult(message)
	return c.Status(fiber.StatusOK).JSON(result)
}

// Error response with data
func ErrorResponse(c *fiber.Ctx, message string) error {
	result := NewErrorDataResult(message)
	return c.Status(fiber.StatusBadRequest).JSON(result)
}

// Error response without data
func ErrorResponseNoData(c *fiber.Ctx, message string) error {
	result := NewErrorResult(message)
	return c.Status(fiber.StatusBadRequest).JSON(result)
}

type SuccessDataResult struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessDataResult(message string, data interface{}) *SuccessDataResult {
	return &SuccessDataResult{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

type ErrorDataResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewErrorDataResult(message string) *ErrorDataResult {
	return &ErrorDataResult{
		Status:  "error",
		Message: message,
	}
}

type SuccessResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewSuccessResult(message string) *SuccessResult {
	return &SuccessResult{
		Status:  "success",
		Message: message,
	}
}

type ErrorResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewErrorResult(message string) *ErrorResult {
	return &ErrorResult{
		Status:  "error",
		Message: message,
	}
}
