package todo

import (
	"backend/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *TodoHandler) GetAllByUser(c echo.Context) error {
	return nil
}

func (h *TodoHandler) GetTodo(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	todo, err := h.todoStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newTodoResponse(todo))
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	req := &todoRequest{}
	todo, err := req.bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if todo, err = h.todoStore.Create(todo); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newTodoResponse(todo))
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	todo, err := h.todoStore.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	req := &todoRequest{}
	new_todo, err := req.bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	todo.Title = new_todo.Description
	todo.Description = new_todo.Description
	todo.Done = new_todo.Done
	todo.Due_at = new_todo.Due_at

	if todo, err = h.todoStore.Update(todo); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newTodoResponse(todo))
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err = h.todoStore.Delete(id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, nil)
}
