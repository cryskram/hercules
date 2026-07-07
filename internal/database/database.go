package database

import (
	"log"
	"time"

	"github.com/cryskram/hercules/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.App.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get db instance:", err)
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to db")

	return db
}
