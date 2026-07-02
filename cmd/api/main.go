package main

import (
	"go-todo2/db"
	"log"
	"net/http"

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
		completed BOOLEAN NOT NULL DEFAULT false
	)
`)
	if err != nil {
		log.Fatal("failed to create todos table:", err)
	}

	log.Println("Connected to the database successfully!")
	log.Println("todos table created successfully!")

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

	// デフォルトではポート 8080 でサーバーを起動します。
	r.Run(":8080")
}
