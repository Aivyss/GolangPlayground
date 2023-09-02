package router

import (
	"com.playground/24_echo/controller"
	"com.playground/24_echo/middleware"
	"github.com/labstack/echo/v4"
)

func CompanyBinding(e *echo.Echo) {
	e.PUT("/company", WrapHandlerFunc2(controller.CompanyController.CreateCompany), middleware.VerifiedUser)
}
