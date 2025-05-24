package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, From EcommerceGO")
	app := fiber.New()

	app.Listen("localhost:3000")
	fmt.Println("Server is running on http://localhost:3000")
}
