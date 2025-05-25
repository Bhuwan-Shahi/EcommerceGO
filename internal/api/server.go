package api

import (
	"ecommerceGO/config"
	"ecommerceGO/internal/api/rest"
	"ecommerceGO/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.Config) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}
	setupRoutes(rh)

	app.Listen(config.Port)
}

func setupRoutes(rh *rest.RestHandler) {
	//users
	handlers.SetupUserRoutes(rh)
	//transcation
	//catalog
}
