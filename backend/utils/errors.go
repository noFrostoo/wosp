package utils

import (
	"github.com/labstack/echo/v4"
	"fmt"
)

// TODO !
type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	fmt.Print("test")

	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}

	return e
}
