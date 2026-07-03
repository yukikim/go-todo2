package service

import (
	"errors"
	"strings"

	"go-todo2/internal/model"
	"go-todo2/internal/repository"
)

// repo は TodoRepository のポインタで、データベース操作を行うためのリポジトリです。
type TodoService struct {
	repo *repository.TodoRepository
}

// NewTodoService はリポジトリを注入した TodoService の新しいインスタンスを作成します。
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
	// タイトルが空でない場合は、リポジトリを使って新しい Todo を作成します。
	return s.repo.Create(title)
}

// GetTodoByID は指定された ID の Todo を取得します。
func (s *TodoService) GetTodoByID(id int) (model.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *TodoService) UpdateTodo(id int, title string, completed bool) (model.Todo, error) {
	// タイトルが空の場合はエラーを返す
	if strings.TrimSpace(title) == "" {
		return model.Todo{}, errors.New("title cannot be empty")
	}

	todo := model.Todo{
		ID:        id,
		Title:     title,
		Completed: completed,
	}

	// repository のUpdate はmodel.Todo1つを受け取り、errorだけを返すので、上のtodoを渡して、エラーがあれば返す
	if err := s.repo.Update(todo); err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}
