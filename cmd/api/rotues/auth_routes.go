package rotues

import (
	"go-rest-api/cmd/api/handlers"
	"go-rest-api/internal/helpers"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, h *handlers.AuthHandler) {
	e.POST("/register", helpers.BindAndValidate(h.Register))
}
