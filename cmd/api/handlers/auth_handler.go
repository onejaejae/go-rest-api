package handlers

import (
	"go-rest-api/cmd/api/requests"
	"go-rest-api/cmd/api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	UserService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
	}
}

func (a *AuthHandler) Register(c echo.Context, request requests.RegisterRequest) error {
	// Check if user already exists
	user, err := a.UserService.GetUserByEmail(request.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
