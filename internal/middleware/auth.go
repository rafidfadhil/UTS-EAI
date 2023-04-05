package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafidfadhil/UTS-EAI/internal/database"
	"github.com/rafidfadhil/UTS-EAI/internal/model"
)

type Config struct {
	Filter       func(*fiber.Ctx) error
	Unauthorized fiber.Handler
}

func UserAuth(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return config.Unauthorized(c)
		}

		// Check for user token
		userToken := model.UserToken{}
		err := database.DB.Where("token = ? AND type = ?", header, "user").First(&userToken).Error
		if err != nil {
			return fiber.ErrUnauthorized
		}

		// User token found, set user info in context and proceed
		c.Locals("userID", userToken.UserID)
		return c.Next()
	}
}

func AdminAuth(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return config.Unauthorized(c)
		}

		// Check for admin token
		adminToken := model.UserToken{}
		err := database.DB.Where("token = ?", header["Authorization"]).First(&adminToken).Error
		if err != nil {
			return config.Unauthorized(c)
		}

		if adminToken.Type != "admin" {
			return config.Unauthorized(c)
        }

		// Admin token found, set user info in context and proceed
		c.Locals("admin", adminToken)
		return c.Next()
	}
}
