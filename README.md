# Go Practice
## Todo API を作成

**「APIの最小仕様」と「DB起動・接続の土台」を作る**

**1. Todo API の最小仕様を決める**

作るエンドポイントを決めます。

```txt
GET    /todos      Todo一覧取得
POST   /todos      Todo作成
GET    /todos/:id  Todo詳細取得
PUT    /todos/:id  Todo更新
DELETE /todos/:id  Todo削除
```

まずはこの2つ。

```txt
GET  /todos
POST /todos
```

**2. Goプロジェクトを作る**

```bash
mkdir go-todo
cd go-todo
go mod init go-todo
```

GinとPostgreSQLドライバを入れます。

```bash
go get github.com/gin-gonic/gin
go get github.com/lib/pq
```

**3. DockerでPostgreSQLを起動できるようにする**

 `docker-compose.yml` 作成

```yaml
services:
  postgres:
    image: postgres:17-alpine
    container_name: go-todo-postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: go_todo
    ports:
      - "5432:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data

volumes:
  postgres-data:
```

起動します。

```bash
docker compose up -d
```

**4. ディレクトリ構成を作る**

```txt
go-todo/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   │   └── todo_handler.go
│   ├── service/
│   │   └── todo_service.go
│   ├── repository/
│   │   └── todo_repository.go
│   ├── model/
│   │   └── todo.go
│   └── db/
│       └── db.go
├── docker-compose.yml
├── go.mod
└── go.sum
```

まずは `main.go`、`model`、`handler` あたりから作る。

**5. 最初に作るファイル**

最初は `cmd/api/main.go` でGinサーバーが起動するところまで作ります。

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
```

起動確認します。

```bash
go run ./cmd/api
```

ブラウザまたはcurlで確認します。

```bash
curl http://localhost:8080/health
```

返ってくればOKです。

```json
{"status":"ok"}
```

**6. 次にDB接続を作る**

`internal/db/db.go` を作って、PostgreSQLにつなげるようにします。

```go
package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/go_todo?sslmode=disable"

	return sql.Open("postgres", dsn)
}
```

この時点で、**Ginのルーティングより先にDB接続確認をしておくこと**が大事です。APIは動いているけどDBに保存できない、という詰まり方を防げます。

**7. Todoテーブルを決める**

最初のテーブルはこれで十分です。

```sql
CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  completed BOOLEAN NOT NULL DEFAULT false
);
```

最初はGo側で起動時に `CREATE TABLE IF NOT EXISTS` を実行してもよいです。慣れてきたら `migrations/` ディレクトリを作って、マイグレーション管理に移る。

**結論**

最初にやることを一言で言うと、

```txt
1. Todo APIの最小仕様を決める
2. go mod init する
3. docker-compose.yml でPostgreSQLを起動する
4. Ginで /health を作ってAPI起動確認する
5. GoからPostgreSQLへ接続確認する
```

です。

最初のゴールは、TodoのCRUD完成ではなく、まずこの状態です。

```txt
Ginサーバーが起動する
PostgreSQLがDockerで起動する
GoからPostgreSQLに接続できる
/health が返る
```

ここまでできたら、次に `model -> repository -> service -> handler` の順でTodo作成処理を足していく。