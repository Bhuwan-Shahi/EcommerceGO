package handlers

import (
	"ecommerceGO/internal/api/rest"
	"ecommerceGO/internal/dto"
	"ecommerceGO/internal/repository"
	"ecommerceGO/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHnadler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	//create an instance of user service and inject ot handler
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
	}
	handler := UserHnadler{
		svc: svc,
	}

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

	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Provide valid input",
		})
	}

	token, err := h.svc.SignUp(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"mesage": "error on signup",
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mesage": token,
	})
}

func (h *UserHnadler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Provide valid input",
		})
	}
	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil{
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Please provide valid username and password",
		})
	}
	
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
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
