package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rafidfadhil/UTS-EAI/internal/database"
	"github.com/rafidfadhil/UTS-EAI/internal/routes"
)

func Init() {
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMODE:  os.Getenv("DB_SSLMODE"),
	}

	database.Connect(config)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
