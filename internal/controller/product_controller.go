package controller

import (
	"time"

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
		"data":    response,
	})

}

func AddProduct(c *fiber.Ctx) (err error) {
	req := dto.ProductRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// Check if category exists
	category := model.Category{}
	err = database.DB.Where("id =?", req.CategoryID).First(&category).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	product := model.Product{
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = database.DB.Create(&product).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Preload Category Name
	err = database.DB.Model(&product).Preload("Category").First(&product).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	response := dto.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Price:        product.Price,
		CategoryName: product.Category.Name,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Successfully Created Product",
		"response": response,
	})
}

// Get All Product
func GetAllProduct(c *fiber.Ctx) (err error) {
	var products []model.Product

	err = database.DB.Find(&products).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Preload Category Name
	err = database.DB.Model(&products).Preload("Category").Find(&products).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	var response []dto.ProductResponse
	for _, product := range products {
		response = append(response, dto.ProductResponse{
			ID:           product.ID,
			Name:         product.Name,
			Price:        product.Price,
			CategoryName: product.Category.Name,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully retrieved all products",
		"data":    response,
	})

}

// Update Product
func UpdateProduct(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	req := dto.ProductRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// Check if category exists

	category := model.Category{}
	err = database.DB.Where("id =?", req.CategoryID).First(&category).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	product := model.Product{}

	err = database.DB.Where("id =?", id).First(&product).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	product.Name = req.Name
	product.Price = req.Price
	product.CategoryID = req.CategoryID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err = database.DB.Save(&product).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// preload category name
	err = database.DB.Model(&product).Preload("Category").First(&product).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	response := dto.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Price:        product.Price,
		CategoryName: product.Category.Name,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Successfully Updated Product",
		"response": response,
	})
}

// Delete Product
func DeleteProduct(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	product := model.Product{}

	err = database.DB.Where("id =?", id).First(&product).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	err = database.DB.Delete(&product).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Successfully Deleted Product",
		"response": nil,
	})
}
