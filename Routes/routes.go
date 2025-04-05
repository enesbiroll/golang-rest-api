package routes

import (
	middleware "rest-api/Middlewares"
	"rest-api/controllers"
	"time"

	"github.com/gofiber/fiber/v2"
)

func StudentRoute(app *fiber.App) {
	// Route: Öğrenci oluşturma (10 saniyede max 5 istek, 2 dakika ban)
	app.Post("/students",
		middleware.RateLimitMiddleware(
			5,              // max istek sayısı
			10*time.Second, // süre penceresi
			2*time.Minute,  // ban süresi
			"10 saniyede 5'ten fazla istek attınız, 2 dakika banlandınız."),
		controllers.CreateStudent)

	// Route: Tüm öğrencileri getirme (5 saniyede max 10 istek, 1 dakika ban)
	app.Get("/students",
		middleware.RateLimitMiddleware(
			10,
			5*time.Second,
			1*time.Minute,
			"Çok fazla istek gönderdiniz. Lütfen 1 dakika bekleyin."),
		controllers.GetStudents)

	// Route: Tek öğrenci getirme
	app.Get("/students/:id", controllers.GetStudentByID)

	// Route: Öğrenci güncelleme
	app.Put("/students/:id", controllers.UpdateStudent)

	// Route: Öğrenci silme
	app.Delete("/students/:id", controllers.DeleteStudent)
}
