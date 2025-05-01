package rotues

import (
	"go-rest-api/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterAppRoutes(e *echo.Echo, h *handlers.AppHandler) {
	e.GET("/", h.HealthCheck)
}
