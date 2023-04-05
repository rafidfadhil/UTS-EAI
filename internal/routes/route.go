package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafidfadhil/UTS-EAI/internal/controller"
	"github.com/rafidfadhil/UTS-EAI/internal/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	// ============== Auth ==============
	auth := api.Group("/auth")

	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)

	// ====================================

	// ============== Categories =============
	categories := api.Group("/categories").Use(middleware.AdminAuth(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		},
	}))

	categories.Post("/create", controller.CreateCategory)
	categories.Get("/", controller.GetAllCategory)

	// =========================================

	// ============== Product =============
	products := api.Group("/product").Use(middleware.AdminAuth(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		},
	}))

	products.Post("/create", controller.AddProduct)
	products.Put("/update/:id", controller.UpdateProduct)
	products.Delete("/delete/:id", controller.DeleteProduct)

	// =========================================

	// ============== Global =============
	global := api.Group("/global")

	// get all products
	global.Get("/products", controller.GetAllProduct)

}
