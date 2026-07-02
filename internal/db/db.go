package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/go_todo?sslmode=disable"

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
