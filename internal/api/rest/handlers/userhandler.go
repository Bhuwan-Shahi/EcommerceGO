package handlers

import (
	"ecommerceGO/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHnadler struct {
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	//create an instance of user service and inject ot handler
	handler := UserHnadler{}

	//pu8blic endpointss
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	//private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("verify", handler.Verify)
	app.Post("profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("order", handler.GetOrders)
	app.Get("order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.Login)

}

func (h *UserHnadler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Register",
	})
}
func (h *UserHnadler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Register",
	})
}

func (h *UserHnadler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Get verification code",
	})
}

func (h *UserHnadler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Verify",
	})
}

func (h *UserHnadler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Crate profile",
	})
}

func (h *UserHnadler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Get profile",
	})
}

func (h *UserHnadler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Add to cart",
	})
}
func (h *UserHnadler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Get cart",
	})
}

func (h *UserHnadler) CrateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Create order",
	})
}

func (h *UserHnadler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Get order",
	})
}
func (h *UserHnadler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Get Order by id",
	})
}

func (h *UserHnadler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": "Become seller",
	})
}
