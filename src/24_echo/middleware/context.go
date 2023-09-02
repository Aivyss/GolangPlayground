package middleware

import (
	"com.playground/jwt"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	claimData jwt.ClaimData
}

func (c CustomContext) ClaimData() *jwt.ClaimData {
	return &c.claimData
}

func NewCustomContext(e echo.Context, data jwt.ClaimData) *CustomContext {
	return &CustomContext{
		Context:   e,
		claimData: data,
	}
}
