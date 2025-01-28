package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kostjaog/go-to-do/pkg/api"
	"github.com/kostjaog/go-to-do/pkg/config"
	"github.com/kostjaog/go-to-do/pkg/model"
	"github.com/kostjaog/go-to-do/pkg/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Формируем строку подключения
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)

	var err error
	// Подключение к базе данных
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных: ", err)
	}

	// Автоматическое создание таблиц
	if err := db.AutoMigrate(&model.Todo{}).Error; err != nil {
		log.Fatal("Ошибка при миграции базы данных: ", err)
	}
}

func main() {
	// Создание репозитория
	todoRepo := repository.NewTodoRepository(db)

	// Настройка маршрутов
	r := api.SetupRouter(todoRepo)

	// Загружаем кастомный порт
	cfg := config.LoadConfig()

	// Запуск сервера на кастомном порту
	fmt.Printf("Запуск сервера. Порт: %s...\n", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
