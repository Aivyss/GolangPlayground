package middleware

import (
	"com.playground/jwt"
	"github.com/labstack/echo/v4"
)

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessTokenCookie, err := c.Cookie("access")

			if err != nil {
				return next(NewCustomContext(c, jwt.ClaimData{
					Verified: false,
				}))
			} else {
				data, err := jwt.ParseAccessTokenString(accessTokenCookie.Value)

				if err != nil {
					return next(NewCustomContext(c, jwt.ClaimData{
						Verified: false,
					}))
				} else {
					return next(NewCustomContext(c, data))
				}
			}
		}
	})
}
