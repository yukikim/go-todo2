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
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, title, completed", title, false,
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

// FindByID は指定された ID の Todo を取得します。見つからなかった場合は空の Todo を返します。
func (r *TodoRepository) FindByID(id int) (model.Todo, error) {
	var todo model.Todo
	err := r.db.QueryRow(`SELECT id, title, completed FROM todos WHERE id = $1`, id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Todo{}, nil // Todo が見つからなかった場合は空の Todo を返す
		}
		return model.Todo{}, err
	}
	return todo, nil
}

// Update は指定された Todo を更新します。
func (r *TodoRepository) Update(todo model.Todo) error {
	_, err := r.db.Exec(`UPDATE todos SET title = $1, completed = $2 WHERE id = $3`, todo.Title, todo.Completed, todo.ID)
	return err
}

// Delete は指定された ID の Todo を削除します。
func (r *TodoRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM todos WHERE id = $1`, id)
	return err
}
