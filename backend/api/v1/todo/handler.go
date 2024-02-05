package todo

import (
	"errors"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoStore Store
	jwtSecret string
}

func NewHandler(todoS Store) (*TodoHandler, error) {
	secret, ok := os.LookupEnv("Signing_Key")
	if !ok {
		return nil, errors.New("No secret key ")
	}

	return &TodoHandler{
		todoStore: todoS,
		jwtSecret: secret,
	}, nil
}

func (h *TodoHandler) Register(group *echo.Group) {
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(h.jwtSecret),
	})

	todo := group.Group("/todo", jwtMiddleware)
	todo.POST("", h.CreateTodo)
	todo.GET("", h.GetAllByUser)
	todo.GET("/:id", h.GetTodo)
	todo.PUT("/:id", h.UpdateTodo)
	todo.DELETE("/:id", h.DeleteTodo)
}
