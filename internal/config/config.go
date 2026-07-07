package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseURL string
}

var App Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env")
	}

	viper.AutomaticEnv()

	App = Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
	}

	if App.Port == "" {
		App.Port = "8080"
	}
}
