package controller

import "github.com/gin-gonic/gin"

type CustomHandlerFunc[T any] func(success bool, req *T, c *gin.Context)

func (r *CustomHandlerFunc[T]) Execute(c *gin.Context) {
	var dto T
	err := c.ShouldBind(&dto)

	(*r)(err == nil, &dto, c)
}

func WrapHandlerFunc[T any](customHandlerFunc func(success bool, req *T, c *gin.Context)) func(c *gin.Context) {
	c := CustomHandlerFunc[T](customHandlerFunc)

	return (&c).Execute
}
