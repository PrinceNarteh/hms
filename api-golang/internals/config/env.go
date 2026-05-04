package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

func initConfig() *config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load .env file")
	}

	return &config{
		AppConfig: &appConfig{
			Port:      fmt.Sprintf(":%s", getEnv("APP_PORT", "4000")),
			RateLimit: getEnvInt("RATE_LIMIT", 10),
		},
		DBConfig: &dbConfig{
			Name: getEnv("DB_NAME", "hsm-db"),
			URI:  getEnv("DB_URI", "mongodb://localhost:27017/hms-db"),
		},
		JwtConfig: &jwtConfig{
			Secret:    getEnv("JWT_SECRET", "secret_password"),
			ExpiresIn: getEnv("EXPIRES_IN", "10m"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value := getEnv(key, "")
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valueInt
}
