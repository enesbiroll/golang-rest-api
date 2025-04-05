package controllers

import (
	"rest-api/Models"
	"rest-api/core/logger"
	"rest-api/services"
	"rest-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// CreateStudent handles the creation of a new student
func CreateStudent(c *fiber.Ctx) error {
	var student Models.Student

	if err := c.BodyParser(&student); err != nil {
		logger.Log.WithField("error", err).Error("Failed to parse request body")
		return utils.ErrorResponse(c, "Invalid input data")
	}

	if err := services.CreateStudent(&student); err != nil {
		logger.Log.WithField("error", err).Error("Failed to create student")
		return utils.ErrorResponse(c, "Failed to create student")
	}

	logger.Log.WithFields(logrus.Fields{
		"student_id":   student.Id,
		"student_name": student.Name,
	}).Info("Student created successfully")

	return utils.SuccessResponse(c, "Student created successfully", student)
}

// GetStudents retrieves all students from the database
func GetStudents(c *fiber.Ctx) error {
	students, err := services.GetAllStudents()
	if err != nil {
		logger.Log.WithField("error", err).Error("Failed to fetch students")
		return utils.ErrorResponse(c, "Failed to fetch students")
	}

	logger.Log.Info("Students retrieved successfully")
	return utils.SuccessResponse(c, "Students retrieved successfully", students)
}

// GetStudentByID retrieves a student by their ID
func GetStudentByID(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := services.GetStudentByID(id)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"student_id": id,
			"error":      err,
		}).Error("Student not found")
		return utils.ErrorResponse(c, "Student not found")
	}

	logger.Log.WithFields(logrus.Fields{
		"student_id":   student.Id,
		"student_name": student.Name,
	}).Info("Student found")

	return utils.SuccessResponse(c, "Student found", student)
}

// UpdateStudent updates an existing student's details
func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedStudent Models.Student
	if err := c.BodyParser(&updatedStudent); err != nil {
		logger.Log.WithField("error", err).Error("Failed to parse request body")
		return utils.ErrorResponse(c, "Invalid input data")
	}

	student, err := services.UpdateStudent(id, &updatedStudent)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"student_id": id,
			"error":      err,
		}).Error("Failed to update student")
		return utils.ErrorResponse(c, "Failed to update student")
	}

	logger.Log.WithFields(logrus.Fields{
		"student_id":   student.Id,
		"student_name": student.Name,
	}).Info("Student updated successfully")

	return utils.SuccessResponse(c, "Student updated successfully", student)
}

// DeleteStudent performs a soft delete of a student
func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := services.DeleteStudent(id); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"student_id": id,
			"error":      err,
		}).Error("Failed to delete student")
		return utils.ErrorResponse(c, "Failed to delete student")
	}

	logger.Log.WithField("student_id", id).Info("Student deleted successfully")
	return utils.SuccessResponseNoData(c, "Student deleted successfully")
}
