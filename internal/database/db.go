package database

import (
	"log"

	"github.com/rafidfadhil/UTS-EAI/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	err = conn.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.UserToken{},
		&model.Category{},
		&model.Product{},
	)

	if err != nil {
		log.Fatal(err)
	}

}
