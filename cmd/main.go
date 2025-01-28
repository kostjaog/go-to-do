package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kostjaog/go-to-do/config"
	_ "github.com/kostjaog/go-to-do/docs"
	"github.com/kostjaog/go-to-do/internal/todo/api"
	"github.com/kostjaog/go-to-do/internal/todo/model"
	"github.com/kostjaog/go-to-do/internal/todo/repository"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)

	var err error
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных: ", err)
	}

	if err := db.AutoMigrate(&model.Todo{}).Error; err != nil {
		log.Fatal("Ошибка при миграции базы данных: ", err)
	}
}

func main() {
	// Создаем центральный маршрутизатор
	r := mux.NewRouter()

	// Добавляем маршруты Swagger
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Создаем репозиторий todo
	todoRepo := repository.NewTodoRepository(db)

	// Регистрируем маршруты для todo
	api.RegisterTodoRoutes(r, todoRepo)

	cfg := config.LoadConfig()

	fmt.Printf("Запуск сервера. Порт: %s...\n", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
