package controller

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: string(hashedPassword),
		Role:     "user",
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

// Login
func Login(c *fiber.Ctx) (err error) {
	req := dto.LoginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Body Request!",
		})
	}

	user := model.User{}

	err = database.DB.Where("email =?", req.Email).First(&user).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to login!",
		})
	}

	// Compare Password with hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Incorrect Password!",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	userToken := model.UserToken{
		UserID: user.ID,
		Token:  token,
		Type:   user.Role,
	}

	// create user token

	err = database.DB.Create(&userToken).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	response := dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  user.Role,
		Token: token,
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User logged in successfully!",
		"data":    response,
	})

}
