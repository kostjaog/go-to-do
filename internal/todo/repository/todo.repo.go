package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kostjaog/go-to-do/internal/todo/model"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) FindOne(id string) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) DeleteOne(id string) (*model.Todo, error) {
	todo, err := r.FindOne(id)
	if err != nil {
		return nil, err
	}
	// Удаляем найденную задачу
	if err := r.db.Delete(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
