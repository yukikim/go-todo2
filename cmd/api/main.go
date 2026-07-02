package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
