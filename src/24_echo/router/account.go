package router

import (
	"com.playground/24_echo/controller"
	"github.com/labstack/echo/v4"
)

func AccountBinding(e *echo.Echo) {
	e.POST("/login", WrapHandlerFunc(controller.AccountController.Login))
	e.POST("/signup", WrapHandlerFunc(controller.AccountController.Signup))
}
