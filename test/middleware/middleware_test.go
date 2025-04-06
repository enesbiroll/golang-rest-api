package middleware_test

import (
	"bytes"
	"net/http"
	middleware "rest-api/middlewares"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitMiddleware(t *testing.T) {
	app := fiber.New()

	// Rate limit middleware'i ile bir rota oluştur
	app.Post("/test", middleware.RateLimitMiddleware(5, 10*time.Second, 2*time.Minute, "Too many requests"), func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// İlk isteği gönder
	resp, err := app.Test(fiberRequest("POST", "/test", ""))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode) // İlk istek başarılı

	// İkinci isteği hemen gönder (yine başarılı)
	resp, err = app.Test(fiberRequest("POST", "/test", ""))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// 3. istek için 1 saniye bekle
	time.Sleep(1 * time.Second)

	// Üçüncü isteği gönder
	resp, err = app.Test(fiberRequest("POST", "/test", ""))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// 4. istek için 1 saniye bekle
	time.Sleep(1 * time.Second)

	// Dördüncü isteği gönder
	resp, err = app.Test(fiberRequest("POST", "/test", ""))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// 5. istek için 1 saniye bekle
	time.Sleep(1 * time.Second)

	// Beşinci isteği gönder
	resp, err = app.Test(fiberRequest("POST", "/test", ""))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// 6. istek oran sınırını geçiyor, bu yüzden 429 bekleniyor
	resp, err = app.Test(fiberRequest("POST", "/test", ""))
	assert.Nil(t, err)
	assert.Equal(t, 429, resp.StatusCode) // 429 bekleniyor
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
