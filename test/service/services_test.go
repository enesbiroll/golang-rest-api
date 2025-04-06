package services_test

import (
	"rest-api/config"
	"rest-api/core/logger" // Logger'ı import et
	"rest-api/models"
	"rest-api/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	// Logger'ı başlat
	logger.Init()

	// Veritabanı bağlantısını başlat
	config.Connect()

	// Öğrenci verisini oluştur
	student := &models.Student{
		Name:        "Test Student",
		StudentCode: "S12345",
	}

	// Servisi çağırarak öğrenci oluştur
	err := services.CreateStudent(student)

	// Testin geçerli olup olmadığını kontrol et
	assert.Nil(t, err)
	assert.NotNil(t, student.Id)
}
