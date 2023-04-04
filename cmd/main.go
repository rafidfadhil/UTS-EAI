package main

import (
	"log"
	
	"github.com/rafidfadhil/UTS-EAI/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// read .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	app.Init()
}
