package services

import (
	"rest-api/Config"
	"rest-api/Models"
	"time"
)

func CreateStudent(student *Models.Student) error {
	return Config.DB.Create(student).Error
}

func GetAllStudents() ([]Models.Student, error) {
	var students []Models.Student
	err := Config.DB.Find(&students).Error
	return students, err
}

func GetStudentByID(id string) (*Models.Student, error) {
	var student Models.Student
	if err := Config.DB.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func UpdateStudent(id string, updatedData *Models.Student) (*Models.Student, error) {
	var student Models.Student
	if err := Config.DB.First(&student, id).Error; err != nil {
		return nil, err
	}

	student.Name = updatedData.Name
	student.StudentCode = updatedData.StudentCode
	if err := Config.DB.Save(&student).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func DeleteStudent(id string) error {
	var student Models.Student
	if err := Config.DB.First(&student, id).Error; err != nil {
		return err
	}

	now := time.Now()
	student.DeletedAt = &now

	return Config.DB.Save(&student).Error
}
