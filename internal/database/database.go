package database

import (
	"log"

	"github.com/cryskram/hercules/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.App.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}

	log.Println("Connected to db")

	return db
}
