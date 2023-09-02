package controller

import (
	"com.playground/24_echo/middleware"
	"com.playground/dto"
	"com.playground/response"
	"net/http"
)

var CompanyController = struct {
	CreateCompany func(successBind bool, req *dto.CreateCompany, c middleware.CustomContext)
}{
	CreateCompany: func(successBind bool, req *dto.CreateCompany, c middleware.CustomContext) {
		_ = c.JSON(http.StatusOK, response.Response[bool]{
			Success: true,
			Status:  http.StatusOK,
			Data:    true,
		})
	},
}
