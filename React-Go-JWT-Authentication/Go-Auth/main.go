package main

import (
	database "Go-Auth/Database"
	pageroutes "Go-Auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
	}))
	pageroutes.RegisterRoutes(app)
	app.Listen(":8080") // Starts the server on port 3000
}
