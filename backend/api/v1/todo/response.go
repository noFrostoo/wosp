package todo

import (
	"backend/models"

	"github.com/google/uuid"
)

type todo struct {
	Id          uuid.UUID `json:"id"`
	User_id     uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Due_at      string    `json:"due_at"`
}

type todoResponse struct {
	Todo todo `json:"todo"`
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

type todosResponse struct {
	Todos []todo `json:"todos"`
}

func newTodosResponse(todos *[]models.Todo) *todosResponse {
	r := &todosResponse{}
	r.Todos = make([]todo, 0)
	for _, t := range *todos {
		r.Todos = append(r.Todos, todo{
			Id:          t.Id,
			User_id:     t.User_id,
			Title:       t.Title,
			Description: t.Description,
			Done:        t.Done,
			Due_at:      t.Due_at,
		})
	}

	return r
}
