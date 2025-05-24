package api

import (
	"ecommerceGO/config"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.Config) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	app.Listen("localhost:8080")
}

func HealthCheck(c *fiber.Ctx) error {
	log.Println("Health check endpoint hit")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Server is running",
		"data":    nil,
	})
}
