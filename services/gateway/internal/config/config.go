package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
	Env string
	AuthUrl string
}

var Cfg = &Config{}

/**
* @description
	If ENV=dev - use dev.env
	Else - use container env variables
*/
func LoadConfig(path string) (*Config, error) {
	var err error
	env := os.Getenv("ENV")
	if os.Getenv("ENV") == "" {
		log.Fatal("ENV variable is not set")
	}
	fmt.Println("WORKING IN ENV :>>", env)

	// use .env file if ENV=dev
	if env == "dev" {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Failed to load env variables :>> %v", err)
	}
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal("ENV PORT is not numeric: ", err)
		}
		if os.Getenv("AUTH_URL") == "" {
			log.Fatal("ENV AUTH_URL is not set")
		}

		Cfg = &Config{
			Port: port,
			Env: env,
			AuthUrl: os.Getenv("AUTH_URL"),
		}

	return Cfg, err
}

func GetConfig() *Config {
	return Cfg
}