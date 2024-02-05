package todo

import (
	"backend/models"

	"github.com/google/uuid"
)


type todoResponse struct {
	Todo struct {
		Id 			uuid.UUID `json:"id"`
		User_id     uuid.UUID `json:"user_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Done        bool      `json:"done"`
		Due_at      string    `json:"due_at"`
	} `json:"user"`
}

func newTodoResponse(todo *models.Todo) todoResponse {
	t := todoResponse{}
	t.Todo.Id = todo.Id
	t.Todo.Title = todo.Title
	t.Todo.Description = todo.Description
	t.Todo.Done = todo.Done
	t.Todo.Due_at = todo.Due_at

	return t
}