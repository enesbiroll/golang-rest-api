package config

import (
	"fmt"
	"rest-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Manuel olarak veritabanı bağlantı bilgilerini giriyoruz
	dbhost := "localhost"
	dbuser := "root"
	dbpass := ""
	dbname := "golangtest"

	// Bağlantı dizesini oluştur
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbname)

	// Veritabanı bağlantısını aç
	dbConnection, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Connection failed to the database") // Hata alırsanız, panik yapın
	}

	// DB nesnesine atayın
	DB = dbConnection
	fmt.Println("DB connected successfully")

	// Veritabanı model göçü işlemi
	AutoMigrate(dbConnection)
}

func AutoMigrate(connection *gorm.DB) {
	// Veritabanı modellerini otomatik olarak göç ettir
	connection.Debug().AutoMigrate(&models.Student{}, &models.Log{}, &models.User{})
	fmt.Println("Database Migrated Successfully")
}
