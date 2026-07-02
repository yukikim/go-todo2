package repository

import (
	"database/sql"

	"go-todo2/internal/model"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// Create は新しい Todo を作成し、作成された Todo を返します。
func (r *TodoRepository) Create(title string) (model.Todo, error) {
	// SQL クエリを実行して、新しい Todo を作成します。
	var todo model.Todo
	err := r.db.QueryRow(
		"INSERT INTO todos (title, completed) VALUES ($1) RETURNING id, title, completed", title,
	).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) FindAll() ([]model.Todo, error) {
	// SQL クエリを実行して、すべての Todo を取得します。
	rows, err := r.db.Query(`SELECT id, title, completed FROM todos ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}

	// rows.Next() は、次の行が存在するかどうかを確認します。存在する場合は true を返し、次の行に移動します。
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil

}
