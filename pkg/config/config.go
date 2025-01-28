package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	AppPort string
}

func LoadConfig() *Config {
	// Загружаем переменные из .env
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env, используем системные переменные")
	}

	// Читаем переменные окружения
	cfg := &Config{
		DBHost:     getEnv("DB_HOST"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
		DBPort:     getEnv("DB_PORT"),

		AppPort: getEnv("APP_PORT"),
	}

	// Проверяем наличие всех значений
	cfg.validate()
	return cfg
}

func getEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Переменная окружения %s не задана", key)
	}
	return val
}

func (cfg *Config) validate() {
	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" || cfg.DBPort == "" || cfg.AppPort == "" {
		log.Fatalf("Некоторые переменные окружения не заданы")
	}
}
