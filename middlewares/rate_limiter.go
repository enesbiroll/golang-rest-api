package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type clientData struct {
	Count     int
	FirstSeen time.Time
}

var (
	requestsMap = make(map[string]*clientData)
	bannedIPs   = make(map[string]time.Time)
	mutex       sync.Mutex
)

// Rate limit middleware
func RateLimitMiddleware(limitCount int, limitDuration time.Duration, banDuration time.Duration, customMessage string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		now := time.Now()

		mutex.Lock()
		defer mutex.Unlock()

		// IP banlı mı kontrol et
		if banTime, banned := bannedIPs[ip]; banned {
			if now.Sub(banTime) < banDuration {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"status":  "error",
					"message": customMessage,
				})
			}
			delete(bannedIPs, ip)
		}

		client, exists := requestsMap[ip]

		// Eğer IP'nin verisi yoksa, yeni bir veri oluştur
		if !exists || now.Sub(client.FirstSeen) > limitDuration {
			// Yeni veriyi başlat
			requestsMap[ip] = &clientData{Count: 1, FirstSeen: now}
		} else {
			client.Count++
			if client.Count > limitCount {
				// Oran sınırlama tetiklendi, IP'yi yasakla
				bannedIPs[ip] = now
				delete(requestsMap, ip)
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"status":  "error",
					"message": customMessage,
				})
			}
		}

		return c.Next()
	}
}
