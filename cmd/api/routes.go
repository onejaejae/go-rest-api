package main

func (app *Application) routes() {
	app.server.GET("/", app.handler.HealthCheck)
}
