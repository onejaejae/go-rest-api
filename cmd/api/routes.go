package main

import (
	"go-rest-api/internal/helpers"
)

func (app *Application) routes() {
	app.server.GET("/", app.handler.HealthCheck)
	app.server.POST("/register", helpers.BindAndValidate(app.handler.Register))
}
