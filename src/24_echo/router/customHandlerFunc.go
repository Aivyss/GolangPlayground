package router

import (
	"github.com/labstack/echo/v4"
)

type CustomHandlerFunc[T any] func(success bool, req *T, c echo.Context)

func (r *CustomHandlerFunc[T]) Execute(c echo.Context) error {
	dto := new(T)
	err := c.Bind(dto)

	(*r)(err == nil, dto, c)

	return err
}

func WrapHandlerFunc[T any](customHandlerFunc func(success bool, req *T, c echo.Context)) func(echo.Context) error {
	handler := CustomHandlerFunc[T](customHandlerFunc)

	return (&handler).Execute
}
