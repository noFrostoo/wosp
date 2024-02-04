package v1

import (
	"backend/api/v1/todo"
	"backend/api/v1/user"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	userHandler    user.UserHandler
	todoStore    todo.Store
}

func NewHandler(userH user.UserHandler, todoS todo.Store) *Handler {
	return &Handler{
		userHandler:    userH,
		todoStore:    todoS,
	}
}

func (h *Handler) Register(group *echo.Group) {
	h.userHandler.Register(group)

}