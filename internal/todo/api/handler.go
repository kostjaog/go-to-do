package api

import (
	"encoding/json"
	"net/http"

	"github.com/kostjaog/go-to-do/internal/todo/model"
	"github.com/kostjaog/go-to-do/internal/todo/repository"
)

type Handler struct {
	todoRepo *repository.TodoRepository
}

func NewHandler(todoRepo *repository.TodoRepository) *Handler {
	return &Handler{todoRepo: todoRepo}
}

// GetTodos godoc
// @Summary Получить все задачи
// @Description Возвращает список всех задач
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Todo
// @Failure 500 {string} string "Ошибка при получении задач"
// @Router /todos [get]
func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoRepo.GetAll()
	if err != nil {
		http.Error(w, "Не удалось получить задачи", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo godoc
// @Summary Создать новую задачу
// @Description Создает задачу с переданными данными
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body model.Todo true "Задача"
// @Success 200 {object} model.Todo
// @Failure 400 {string} string "Неверный формат данных"
// @Failure 500 {string} string "Не удалось создать задачу"
// @Router /todos [post]
func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}
	if err := h.todoRepo.Create(&todo); err != nil {
		http.Error(w, "Не удалось создать задачу", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}
