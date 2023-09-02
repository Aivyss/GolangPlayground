package router

import (
	"com.playground/24_echo/middleware"
	"github.com/labstack/echo/v4"
)

type CustomHandlerFunc[T any] func(success bool, req *T, c echo.Context)

func (r *CustomHandlerFunc[T]) Execute(c echo.Context) error {
	dto := new(T)
	err := c.Bind(dto)

	(*r)(err == nil, dto, c)

	return err
}

type CustomHandlerFunc2[T any] func(success bool, req *T, c middleware.CustomContext)

func (r *CustomHandlerFunc2[T]) Execute(c echo.Context) error {
	dto := new(T)
	err := c.Bind(dto)

	(*r)(err == nil, dto, *c.(*middleware.CustomContext))

	return err
}

func WrapHandlerFunc[T any](customHandlerFunc func(success bool, req *T, c echo.Context)) func(echo.Context) error {
	handler := CustomHandlerFunc[T](customHandlerFunc)

	return (&handler).Execute
}

func WrapHandlerFunc2[T any](customHandlerFunc func(success bool, req *T, c middleware.CustomContext)) func(echo.Context) error {
	handler := CustomHandlerFunc2[T](customHandlerFunc)

	return (&handler).Execute
}
