package handlers

import (
	"ecommerceGO/internal/api/rest"
	"ecommerceGO/internal/dto"
	"ecommerceGO/internal/repository"
	"ecommerceGO/internal/service"
	"log"
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
		Repo:   repository.NewUserRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := UserHnadler{
		svc: svc,
	}
	pubRoutes := app.Group("/users")

	//pu8blic endpointss
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	//private endpoints
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("verify", handler.Verify)
	pvtRoutes.Post("profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)

	pvtRoutes.Post("cart", handler.AddToCart)
	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Get("order", handler.GetOrders)
	pvtRoutes.Get("order/:id", handler.GetOrder)

	pvtRoutes.Post("/become-seller", handler.Login)

}

func (h *UserHnadler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}

	// Parse the request body
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Provide valid input",
		})
	}

	// Call the SignUp method
	token, err := h.svc.SignUp(user)
	if err != nil {
		if err.Error() == "user already exists" {
			return ctx.Status(http.StatusConflict).JSON(&fiber.Map{
				"message": "User with this email already exists",
			})
		}

		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
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

	if err != nil {
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

	user := h.svc.Auth.GetCurrentUser(ctx)

	//create verification code and update to user profile in DB

	code, err := h.svc.GetVerificationCode(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get verification code",
		"data":    code,
	})
}

func (h *UserHnadler) Verify(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	// request
	var req dto.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}
func (h *UserHnadler) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Crate profile",
	})

}

func (h *UserHnadler) GetProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	log.Println(user)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get profile",
		"user":    user,
	})
}

func (h *UserHnadler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Add to cart",
	})
}
func (h *UserHnadler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get cart",
	})
}

func (h *UserHnadler) CrateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Create order",
	})
}

func (h *UserHnadler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get order",
	})
}
func (h *UserHnadler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get Order by id",
	})
}

func (h *UserHnadler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Become seller",
	})
}
