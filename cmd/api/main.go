package main

import (
	"fmt"
	"go-rest-api/cmd/api/handlers"
	"go-rest-api/common"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	e := echo.New()
	err := godotenv.Load()
	db, dbErr := common.NewMySql()

	if dbErr != nil {
		e.Logger.Fatal(dbErr)
	}

	if err != nil {
		e.Logger.Fatal(err)
	}

	appPort := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf(":%s", appPort)

	handler := handlers.Handler{
		DB: db,
	}

	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: handler,
	}

	e.GET("/", app.handler.HealthCheck)

	e.Logger.Fatal(e.Start(appAddress))
}
