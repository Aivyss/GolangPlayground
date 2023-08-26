package router

import (
	"com.playground/23_gin/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	controller.AccountBinding(router)

	return router
}
