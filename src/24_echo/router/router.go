package router

import (
	"com.playground/24_echo/middleware"
	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	e := echo.New()

	AccountBinding(e)
	CompanyBinding(e)
	middleware.RegisterMiddlewares(e)

	return e
}
