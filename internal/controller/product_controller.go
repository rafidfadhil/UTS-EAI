package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafidfadhil/UTS-EAI/internal/database"
	"github.com/rafidfadhil/UTS-EAI/internal/dto"
	"github.com/rafidfadhil/UTS-EAI/internal/model"
)

func CreateCategory(c *fiber.Ctx) (err error) {

	req := dto.CategoryRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	category := model.Category{
		Name: req.Name,
	}

	err = database.DB.Create(&category).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	response := dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"messaage": "Successfully Created Category",
		"response": response,
	})

}

func GetAllCategory(c *fiber.Ctx) (err error) {
	var categories []model.Category

	err = database.DB.Find(&categories).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	var response []dto.CategoryResponse
	for _, category := range categories {
		response = append(response, dto.CategoryResponse{
            ID:   category.ID,
            Name: category.Name,
        })
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Successfully retrieved all categories",
		"data": response,
    }) 

		

}
