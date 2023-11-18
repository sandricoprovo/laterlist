package main

import (
	"log"
	routes "read-it-later/api/routes/link"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

	// Create app
	app := fiber.New()

	// Add routers
	routes.AddLinkRouter(app)

	// app.Get("/ping", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "ping",
	// 	})
	// })

	app.Listen(":8080")
}
