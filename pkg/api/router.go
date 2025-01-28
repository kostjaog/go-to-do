package api

import (
	"github.com/gorilla/mux"
	"github.com/kostjaog/go-to-do/pkg/repository"
)

// RegisterTodoRoutes регистрирует маршруты для todo
func RegisterTodoRoutes(r *mux.Router, repo *repository.TodoRepository) {
	handler := NewHandler(repo)
	r.HandleFunc("/todos", handler.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
}
