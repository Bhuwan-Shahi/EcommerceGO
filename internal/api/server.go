package api

import (
	"ecommerceGO/config"
	"ecommerceGO/internal/api/rest"
	"ecommerceGO/internal/api/rest/handlers"
	"ecommerceGO/internal/domain"
	"ecommerceGO/internal/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.Config) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database Connection Error %v\n", err)
	}
	log.Println("Database connected!!")

	//Running data migration

	db.AutoMigrate(&domain.User{})
	auth := helper.SetupAuth(config.AppSecret)
	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
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
