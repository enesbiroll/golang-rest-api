package controllers

import (
	"rest-api/core/logger"
	Models "rest-api/models"
	"rest-api/services"
	"rest-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// CreateStudent handles the creation of a new student
// @Summary Öğrenci oluşturma
// @Description Yeni bir öğrenci oluşturur
// @Tags Students
// @Accept json
// @Produce json
// @Param student body models.Student true "Öğrenci Bilgileri"
// @Success 200 {object} models.Student "Student created successfully"
// @Failure 400 {string} string "Invalid input data"
// @Failure 500 {string} string "Failed to create student"
// @Router /students [post]
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
// GetStudents retrieves all students from the database
// @Summary Öğrencileri getir
// @Description Tüm öğrencileri getirir
// @Tags Students
// @Accept json
// @Produce json
// @Success 200 {array} models.Student "List of students"
// @Failure 500 {string} string "Failed to fetch students"
// @Router /students [get]
func GetStudents(c *fiber.Ctx) error {
	students, err := services.GetAllStudents()
	if err != nil {
		logger.Log.WithField("error", err).Error("Failed to fetch students")
		return utils.ErrorResponse(c, "Failed to fetch students")
	}
	return utils.SuccessResponse(c, "Students retrieved successfully", students)
}

// GetStudentByID retrieves a student by their ID
// GetStudentByID retrieves a student by their ID
// @Summary Öğrenci bilgilerini getir
// @Description ID'ye göre öğrenci bilgilerini getirir
// @Tags Students
// @Accept json
// @Produce json
// @Param id path string true "Öğrenci ID"
// @Success 200 {object} models.Student "Student found"
// @Failure 400 {string} string "Student not found"
// @Router /students/{id} [get]
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

	return utils.SuccessResponse(c, "Student found", student)
}

// UpdateStudent updates an existing student's details
// UpdateStudent updates an existing student's details
// @Summary Öğrenci bilgilerini güncelle
// @Description Var olan bir öğrencinin bilgilerini günceller
// @Tags Students
// @Accept json
// @Produce json
// @Param id path string true "Öğrenci ID"
// @Param student body models.Student true "Updated student data"
// @Success 200 {object} models.Student "Student updated successfully"
// @Failure 400 {string} string "Invalid input data"
// @Failure 404 {string} string "Student not found"
// @Router /students/{id} [put]
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
// DeleteStudent performs a soft delete of a student
// @Summary Öğrenci silme
// @Description Öğrenciyi siler (soft delete)
// @Tags Students
// @Accept json
// @Produce json
// @Param id path string true "Öğrenci ID"
// @Success 200 {string} string "Student deleted successfully"
// @Failure 404 {string} string "Student not found"
// @Router /students/{id} [delete]
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
