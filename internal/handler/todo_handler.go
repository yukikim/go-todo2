package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-todo2/internal/service"
)

// TodoHandler は TodoService のポインタを持つ構造体で、HTTP リクエストを処理するためのハンドラです。
type TodoHandler struct {
	service *service.TodoService
}

// NewTodoHandler は TodoHandler の新しいインスタンスを作成します。
func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

type CreateTodoRequest struct {
	Title string `json:"title"`
}

// CreateTodo は新しい Todo を作成するための HTTP ハンドラです。
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest

	// c.ShouldBindJSON は、リクエストの JSON ボディを構造体にバインドします。バインドに失敗した場合は、400 Bad Request を返します。
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TodoService の CreateTodo メソッドを呼び出して、新しい Todo を作成します。
	todo, err := h.service.CreateTodo(req.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 作成された Todo を JSON レスポンスとして返します。ステータスコードは 201 Created です。
	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	// TodoService の GetTodos メソッドを呼び出して、すべての Todo を取得します。
	todos, err := h.service.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 取得した Todo を JSON レスポンスとして返します。ステータスコードは 200 OK です。
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	// URL パラメータから ID を取得します。
	idParam := c.Param("id")

	// ID を整数に変換します。変換に失敗した場合は、400 Bad Request を返します。
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// TodoService の GetTodoByID メソッドを呼び出して、指定された ID の Todo を取得します。
	todo, err := h.service.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Todo が見つからなかった場合は、404 Not Found を返します。
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	// 取得した Todo を JSON レスポンスとして返します。ステータスコードは 200 OK です。
	c.JSON(http.StatusOK, todo)
}
