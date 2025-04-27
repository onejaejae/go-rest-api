package handlers

import (
	"go-rest-api/cmd/api/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context, request requests.RegisterRequest) error {

	return c.JSON(http.StatusOK, request)
}
