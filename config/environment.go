package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Environment struct {
	HostPort string
	DBHost string
	DBName string
	DBUsername string
	DBPassword string
}

func GetEnvironment() Environment {
	godotenv.Load()
	return Environment{
		HostPort:   os.Getenv("HOST_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
	}
}
