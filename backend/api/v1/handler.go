package v1

import (
	"backend/api/v1/todo"
	"backend/api/v1/user"
	"backend/store"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	userHandler    user.UserHandler
	todoHandler    todo.TodoHandler
}

func NewHandler(userS *store.UserStore, todoS *store.TodoStore) (*Handler, error) {
	uh, err := user.NewHandler(userS)
	if err != nil {
		return nil, nil
	}

	th, err := todo.NewHandler(todoS)
	if err != nil {
		return nil, nil
	}

	return &Handler{
		userHandler:   *uh,
		todoHandler:   *th,
	}, nil
}

func (h *Handler) Register(group *echo.Group) {
	h.userHandler.Register(group)
	h.todoHandler.Register(group)
}