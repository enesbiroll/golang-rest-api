package routes_test

import (
	"bytes"
	"net/http"
	"rest-api/config"
	"rest-api/core/logger" // Logger'ı import et
	"rest-api/routes"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudentRoute(t *testing.T) {
	// Logger'ı başlat
	logger.Init()

	// Veritabanı bağlantısını başlat
	config.Connect()

	// Yeni bir Fiber uygulaması oluştur
	app := fiber.New()

	// Rotaları tanımla
	routes.StudentRoute(app)

	// Test verisini JSON formatında oluştur
	payload := `{"name":"Test Student","code":"S12345"}`

	// POST isteği gönder
	resp, err := app.Test(fiberRequest("POST", "/students", payload))

	// Yanıt kontrolü
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func fiberRequest(method, url, body string) *http.Request {
	// JSON verisini byte dizisine dönüştür
	reqBody := []byte(body)

	// HTTP isteği oluştur
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err) // Eğer hata olursa panik at
	}

	// İstek başlıklarını ayarlayın
	req.Header.Set("Content-Type", "application/json")
	return req
}
