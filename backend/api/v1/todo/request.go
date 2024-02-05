package todo

import (
	"backend/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type todoRequest struct {
	Todo struct {
		User_id     uuid.UUID `json:"user_id" validate:"required"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description"`
		Done        bool      `json:"done"`
		Due_at      string    `json:"due_at"`
	} `json:"todo"`
}

func (r *todoRequest) bind(c echo.Context) (*models.Todo, error) {
	if err := c.Bind(r); err != nil {
		return nil, err
	}
	if err := c.Validate(r); err != nil {
		return nil, err
	}

	todo := &models.Todo{}
	todo.User_id = r.Todo.User_id
	todo.Title = r.Todo.Title
	todo.Description = r.Todo.Description
	todo.Done = r.Todo.Done
	todo.Due_at = r.Todo.Due_at

	return todo, nil
}
