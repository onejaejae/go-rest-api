package handlers

import "go-rest-api/cmd/api/services"

type Handler struct {
	AuthHandler *AuthHandler
	AppHandler  *AppHandler
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		AuthHandler: NewAuthHandler(service.UserService),
		AppHandler:  NewAppHandler(),
	}
}
