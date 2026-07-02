package service

import (
	"errors"
	"strings"

	"go-todo2/internal/model"
	"go-todo2/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

// NewTodoService は TodoService の新しいインスタンスを作成します。
func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

// GetTodos はすべての Todo を取得します。
func (s *TodoService) GetTodos() ([]model.Todo, error) {
	return s.repo.FindAll()
}

// CreateTodo は新しい Todo を作成します。
func (s *TodoService) CreateTodo(title string) (model.Todo, error) {
	// タイトルが空の場合はエラーを返す
	if strings.TrimSpace(title) == "" {
		return model.Todo{}, errors.New("title cannot be empty")
	}
	return s.repo.Create(title)
}
