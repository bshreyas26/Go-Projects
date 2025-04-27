package routes

import (
	controllers "Go-Auth/Controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.UserDetails)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/verify", controllers.VerifyCode)
}
