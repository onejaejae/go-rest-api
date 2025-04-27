package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("CustomMiddleware")
		c.Response().Header().Set(echo.HeaderServer, "CustomMiddleware")
		return next(c)
	}
}
