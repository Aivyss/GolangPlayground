package middleware

import (
	"com.playground/response"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func VerifiedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		context, ok := c.(*CustomContext)

		if ok && context.claimData.Verified {
			return next(c)
		} else {
			err := errors.New("not verified user")
			_ = c.JSON(http.StatusBadRequest, response.Response[string]{
				Data:    err.Error(),
				Success: false,
				Status:  http.StatusBadRequest,
			})

			return err
		}
	}
}
