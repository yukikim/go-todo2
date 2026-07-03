package main

import (
	"log"
	"net/http"

	"go-todo2/internal/db"
	"go-todo2/internal/handler"
	"go-todo2/internal/repository"
	"go-todo2/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN NOT NULL DEFAULT FALSE
		)
	`)
	if err != nil {
		log.Fatalf("failed to create todos table: %v", err)
	}

	// TodoRepository、TodoService、TodoHandler のインスタンスを作成します。
	todoRepository := repository.NewTodoRepository(database)
	// TodoService は TodoRepository を使用して、ビジネスロジックを実装します。
	todoService := service.NewTodoService(todoRepository)
	// TodoHandler は TodoService を使用して、HTTP リクエストを処理します。
	todoHandler := handler.NewTodoHandler(todoService)

	r := gin.Default()

	// c は gin.Context のポインタで、HTTP リクエストやレスポンスに関する情報を持っています。
	r.GET("/health", func(c *gin.Context) {

		/* c.JSON は http.StatusOK だったら 200 OK のステータスコードを返し、JSON レスポンスを生成します。
		   http.StatusOK じゃない場合は、適切なステータスコードを返すことができます。
		   H は gin.H 型で、JSON レスポンスを簡単に作成するためのマップです。*/
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// TodoHandler のルートを設定します。GET /todos はすべての Todo を取得し、POST /todos は新しい Todo を作成します。
	r.GET("/todos", todoHandler.GetTodos)
	// POST /todos は新しい Todo を作成するためのエンドポイントです。リクエストボディに JSON 形式でタイトルを送信します。
	r.POST("/todos", todoHandler.CreateTodo)
	// GET /todos/:id は指定された ID の Todo を取得するためのエンドポイントです。URL パラメータから ID を取得します。
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	// PUT /todos/:id は指定された ID の Todo を更新するためのエンドポイントです。URL パラメータから ID を取得します。
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	// DELETE /todos/:id は指定された ID の Todo を削除するためのエンドポイントです。URL パラメータから ID を取得します。
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	log.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
