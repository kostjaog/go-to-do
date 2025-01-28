package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/kostjaog/go-to-do/internal/todo/repository"
)

func RegisterTodoRoutes(r chi.Router, repo *repository.TodoRepository) {
	handler := NewHandler(repo)
	r.Get("/todos", handler.GetTodos)               // Получить все задачи
	r.Post("/todos", handler.CreateTodo)            // Создать задачу
	r.Get("/todos/{id}", handler.GetTodoByID)       // Получить задачу по ID
	r.Delete("/todos/{id}", handler.DeleteTodoById) // Удалить задачу по ID
}
