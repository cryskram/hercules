package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cryskram/hercules/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.Load()

	m, err := migrate.New(
		"file://migrations",
		config.App.DatabaseURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run ./cmd/migrate [up|down|version|force <version>]")
	}

	switch os.Args[1] {

	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("✅ Database migrated successfully")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("⬇ Database rolled back")

	case "version":
		v, dirty, err := m.Version()
		if err != nil {
			if err == migrate.ErrNilVersion {
				fmt.Println("No migrations applied.")
				return
			}
			log.Fatal(err)
		}

		fmt.Printf("Current Version: %d\nDirty: %v\n", v, dirty)
	case "force":

		if len(os.Args) < 3 {
			log.Fatal("Usage: go run ./cmd/migrate force <version>")
		}

		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Invalid migration version")
		}

		if err := m.Force(version); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("✅ Forced migration version to %d\n", version)

	default:
		log.Fatal("Unknown command")
	}
}
