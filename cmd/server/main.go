package main

import (
	"log"

	"github.com/cryskram/hercules/internal/config"
	"github.com/cryskram/hercules/internal/database"
	"github.com/cryskram/hercules/internal/handlers"
	repository "github.com/cryskram/hercules/internal/repositories"
	"github.com/cryskram/hercules/internal/routes"
	"github.com/cryskram/hercules/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	db := database.Connect()
	router := gin.Default()

	bondRepo := repository.NewBondRepository(db)
	bondService := services.NewBondService(bondRepo)
	bondHandler := handlers.NewBondHandler(bondService)

	routes.RegisterRoutes(router, bondHandler)

	log.Println("Hercules is running")
	router.Run(":" + config.App.Port)
}
