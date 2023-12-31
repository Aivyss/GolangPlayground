package controller

import (
	"com.playground/23_gin/service"
	dto2 "com.playground/dto"
	"com.playground/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	accountService *service.AccountService
}

var controller = &AccountController{
	accountService: service.AccountServe(),
}

func AccountBinding(r *gin.Engine) {
	r.POST("/signup", WrapHandlerFunc(controller.Signup))
	r.POST("/login", WrapHandlerFunc(controller.Login))
}

func (a AccountController) Signup(successBind bool, req *dto2.Signup, c *gin.Context) {
	if successBind {
		a.accountService.Signup(req)

		c.JSON(http.StatusOK, gin.H{
			"response": response.Response[bool]{
				Status:  http.StatusOK,
				Success: true,
				Data:    true,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": response.Response[bool]{
				Status:  http.StatusBadRequest,
				Success: false,
				Data:    false,
			},
		})
	}
}

func (a AccountController) Login(successBind bool, req *dto2.Login, c *gin.Context) {
	if successBind {
		account, err := a.accountService.Login(req)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"response": response.Response[string]{
					Status:  http.StatusBadRequest,
					Success: false,
					Data:    err.Error(),
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"response": response.Response[dto2.Account]{
					Status:  http.StatusOK,
					Success: true,
					Data:    account,
				},
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": response.Response[bool]{
				Status:  http.StatusBadRequest,
				Success: false,
				Data:    false,
			},
		})
	}
}
