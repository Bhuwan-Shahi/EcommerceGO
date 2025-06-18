package handlers

import (
	"ecommerceGO/internal/api/rest"
	"ecommerceGO/internal/repository"
	"ecommerceGO/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	//create an instance of user service and inject ot handler
	svc := service.UserService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := CatalogHandler{
		svc: svc,
	}
	pubRoutes := app.Group("/users")

	//pu8blic endpointss

	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProduct)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategory)

	selRoutes := app.Group("/seller")

	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)
	selRoutes.Post("/products", handler.CreateProduct)

	//porduct

	selRoutes.Post("products", handler.CreateProduct)
	selRoutes.Get("products", handler.GetProducts)
	selRoutes.Get("products/:id", handler.GetProduct)
	selRoutes.Put("products/:id", handler.EditProduct)
	selRoutes.Patch("products/:id", handler.UpdateStock)
	selRoutes.Delete("products/:id", handler.DeleteProduct)
}

func (h *CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	// Implement the logic to create categories
	return rest.SuccessMessage(ctx, "Category endpoint", nil)

}

func (h *CatalogHandler) EditCategory(ctx *fiber.Ctx) error {
	// Implement the logic to update a category
	return rest.SuccessMessage(ctx, "Update Category endpoint", nil)
}
func (h *CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	// Implement the logic to delete a category
	return rest.SuccessMessage(ctx, "Delete Category endpoint", nil)
}

func (h *CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	// Implement the logic to create a product
	return rest.SuccessMessage(ctx, "Product endpoint", nil)
}
func (h *CatalogHandler) GetProducts(ctx *fiber.Ctx) error {
	// Implement the logic to get products
	return rest.SuccessMessage(ctx, "Get Products endpoint", nil)
}
func (h *CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	// Implement the logic to get a single product
	return rest.SuccessMessage(ctx, "Get Product endpoint", nil)
}
func (h *CatalogHandler) EditProduct(ctx *fiber.Ctx) error {
	// Implement the logic to update a product
	return rest.SuccessMessage(ctx, "Update Product endpoint", nil)
}
func (h *CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	// Implement the logic to update a product
	return rest.SuccessMessage(ctx, "Update Product endpoint", nil)
}
func (h *CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	// Implement the logic to delete a product
	return rest.SuccessMessage(ctx, "Delete Product endpoint", nil)
}
func (h *CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	// Implement the logic to get categories
	return rest.SuccessMessage(ctx, "Get Categories endpoint", nil)
}
func (h *CatalogHandler) GetCategory(ctx *fiber.Ctx) error {
	// Implement the logic to get a single category
	return rest.SuccessMessage(ctx, "Get Category endpoint", nil)
}
