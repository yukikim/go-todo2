package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=go_todo2 sslmode=disable"

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// データベース接続の確認
	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}

// Close はデータベース接続を閉じます。
// func (db *sql.DB) Close() error {
// 	return db.Close()
// }

// NewDB はデータベース接続を作成し、接続が成功した場合に *sql.DB を返します。
// 接続に失敗した場合はエラーを返します。
// func NewDB() (*sql.DB, error) {
// 	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=go_todo2 sslmode=disable"

// 	// sql.Open はデータベース接続を開きます。PostgreSQL ドライバを使用して、指定された DSN (Data Source Name) を使って接続します。
// 	return sql.Open("postgres", dsn)
// }
