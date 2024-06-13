package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PostgresURI string
	Port        string
}

func NewConfig() *Config {
	_ = godotenv.Load("./build/.env")

	postgresURI := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	port := os.Getenv("PORT")
	return &Config{
		PostgresURI: postgresURI,
		Port:        port,
	}
}
