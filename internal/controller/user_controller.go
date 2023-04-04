package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafidfadhil/UTS-EAI/internal/database"
	"github.com/rafidfadhil/UTS-EAI/internal/dto"
	"github.com/rafidfadhil/UTS-EAI/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) (err error) {
	// request body
	req := dto.RegisterRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Body Request!",
		})
	}

	// Hash the Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register!",
		})
	}

	// Create User Role: by default, all users are registered as "user"
	var userRole model.Role

	err = database.DB.Where("name = ?", "user").First(&userRole).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register!",
		})
	}

	user := model.User{
		Name:    req.Name,
		Email:   req.Email,
		Phone:  req.Phone,
		Password: string(hashedPassword),
		Role: "user",
	}

	// Check if the user already exists
	var CheckUser model.User

	err = database.DB.Where("email = ?", req.Email).First(&CheckUser).Error

	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists!",
		})
	}

	// Create User
	err = database.DB.Create(&user).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register!",
		})
	}

	// response
	res := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  user.Role,
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully!",
		"data":    res,
	})
}
