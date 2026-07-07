package main

import (
	"log"

	"github.com/cryskram/hercules/internal/config"
	"github.com/cryskram/hercules/internal/database"
	"github.com/cryskram/hercules/internal/seeder"
)

func main() {

	config.Load()

	db := database.Connect()

	log.Println("📄 Reading Excel...")

	if err := seeder.SeedBonds(db, "assets/BondMaster.xlsx"); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database seeded successfully")
}
