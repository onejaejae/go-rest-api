package helpers

import (
	"go-rest-api/common"

	"github.com/labstack/echo/v4"
)

func BindAndValidate[T any](next func(c echo.Context, req T) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req T
		if err := c.Bind(&req); err != nil {
			return common.SendBadRequestResponse(c, err.Error())
		}
		if err := c.Validate(&req); err != nil {
			return common.SendBadRequestResponse(c, err.Error())
		}
		return next(c, req)
	}
}
