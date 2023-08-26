package controller

import (
	"com.playground/23_gin/dto"
	"com.playground/23_gin/response"
	"com.playground/23_gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
}

var controller = &AccountController{}

func AccountBinding(r *gin.Engine) {
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
}

func (a AccountController) Signup(c *gin.Context) {
	var signup dto.Signup

	if c.ShouldBind(&signup) == nil {
		service.AccountServe().Signup(&signup)

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

func (a AccountController) Login(c *gin.Context) {
	var login dto.LoginDto

	if c.ShouldBind(&login) == nil {
		account, err := service.AccountServe().Login(&login)

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
				"response": response.Response[dto.Account]{
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
