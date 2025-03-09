package main

import (
	"fiber_api/db"
	"fiber_api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	db.Connect()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3450"))
}
