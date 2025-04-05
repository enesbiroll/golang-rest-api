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

// Geliştirilmiş: limitCount, limitDuration, banDuration, customMessage
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

		if !exists || now.Sub(client.FirstSeen) > limitDuration {
			requestsMap[ip] = &clientData{Count: 1, FirstSeen: now}
		} else {
			client.Count++
			if client.Count > limitCount {
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
