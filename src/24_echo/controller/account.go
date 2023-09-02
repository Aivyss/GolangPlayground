package controller

import (
	"com.playground/24_echo/service"
	"com.playground/dto"
	"com.playground/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

var AccountController = struct {
	Signup func(successBind bool, req *dto.Signup, c echo.Context)
	Login  func(successBind bool, req *dto.Login, c echo.Context)
}{
	Signup: func(successBind bool, req *dto.Signup, c echo.Context) {
		if successBind {
			service.AccountService.Signup(req)

			_ = c.JSON(http.StatusOK, response.Response[bool]{
				Status:  http.StatusOK,
				Success: true,
				Data:    true,
			})
		} else {
			_ = c.JSON(http.StatusBadRequest, response.Response[bool]{
				Status:  http.StatusBadRequest,
				Success: false,
				Data:    false,
			})
		}
	},

	Login: func(successBind bool, req *dto.Login, c echo.Context) {
		if successBind {
			account, err := service.AccountService.Login(req)

			if err != nil {
				_ = c.JSON(http.StatusBadRequest, response.Response[string]{
					Status:  http.StatusBadRequest,
					Success: false,
					Data:    err.Error(),
				})
			} else {
				_ = c.JSON(http.StatusOK, response.Response[dto.Account]{
					Status:  http.StatusOK,
					Success: true,
					Data:    account,
				})
			}
		} else {
			_ = c.JSON(http.StatusBadRequest, response.Response[bool]{
				Status:  http.StatusBadRequest,
				Success: false,
				Data:    false,
			})
		}
	}}
