package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) HealthCheck(c echo.Context) error {
	healthCheckStruct := struct {
		Health bool `json:"health"`
	}{
		Health: true,
	}

	return c.JSON(http.StatusOK, healthCheckStruct)
}
