package main

import (
	"log"

	"github.com/cryskram/hercules/internal/config"
	"github.com/cryskram/hercules/internal/database"
	"github.com/cryskram/hercules/internal/routes"
)

func main() {
	config.Load()

	db := database.Connect()

	router := routes.SetupRoutes(db)

	log.Printf("Hercules API is running on :%s", config.App.Port)

	router.Run(":" + config.App.Port)
}
