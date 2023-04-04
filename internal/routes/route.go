package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafidfadhil/UTS-EAI/internal/controller"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	// ============== Auth ==============
	auth := api.Group("/auth")

	auth.Post("/register", controller.Register)
}
