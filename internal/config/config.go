package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort           string
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	JWTSecret         string
	JWTExpiryMinutes  int
}

func Load() *Config {
	// Load .env file (ignore error if not found — production may use env vars directly)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("JWT_EXPIRY_MINUTES", 15)

	return &Config{
		AppPort:          viper.GetString("APP_PORT"),
		DBHost:           viper.GetString("DB_HOST"),
		DBPort:           viper.GetString("DB_PORT"),
		DBUser:           viper.GetString("DB_USER"),
		DBPassword:       viper.GetString("DB_PASSWORD"),
		DBName:           viper.GetString("DB_NAME"),
		JWTSecret:        viper.GetString("JWT_SECRET"),
		JWTExpiryMinutes: viper.GetInt("JWT_EXPIRY_MINUTES"),
	}
}
