package router

import "github.com/labstack/echo/v4"

func GetRouter() *echo.Echo {
	e := echo.New()

	AccountBinding(e)

	return e
}
