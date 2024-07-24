package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/hcontext"
)

func GetUserAgent() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userAgent := c.Request().UserAgent()
			c.Set(hcontext.UserAgent.String(), userAgent)
			return next(c)
		}
	}
}
