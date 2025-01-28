package api

import (
	"github.com/gorilla/mux"
	"github.com/kostjaog/go-to-do/pkg/repository"
)

func SetupRouter(todoRepo *repository.TodoRepository) *mux.Router {
	handler := NewHandler(todoRepo)
	r := mux.NewRouter()
	r.HandleFunc("/todos", handler.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	return r
}
