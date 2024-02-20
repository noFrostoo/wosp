package v1

import (
	"backend/api/v1/todo"
	"backend/api/v1/user"
	"backend/store"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UserHandler user.UserHandler
	TodoHandler todo.TodoHandler
}

func NewHandler(userS *store.UserStore, todoS *store.TodoStore) (*Handler, error) {
	uh, err := user.NewHandler(userS)
	if err != nil {
		return nil, err
	}

	th, err := todo.NewHandler(todoS)
	if err != nil {
		return nil, err
	}

	return &Handler{
		UserHandler: *uh,
		TodoHandler: *th,
	}, nil
}

func (h *Handler) Register(group *echo.Group) {
	h.UserHandler.Register(group)
	h.TodoHandler.Register(group)
}
