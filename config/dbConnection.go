package Config

import (
	"fmt"
	"os"
	"rest-api/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Load environment variables
	godotenv.Load()
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpass := os.Getenv("MYSQL_PASS")
	dbname := os.Getenv("MYSQL_DB_NAME")

	// Connection string
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbname)
	dbConnection, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("connection failed to the database ")
	}
	DB = dbConnection
	fmt.Println("db connected successfully")

	// Auto migrate the models
	AutoMigrate(dbConnection)
}

func AutoMigrate(connection *gorm.DB) {
	// Migrate models, including the Log model
	connection.Debug().AutoMigrate(&Models.Student{}, &Models.Log{})
	fmt.Println("Database Migrated Successfully")
}
