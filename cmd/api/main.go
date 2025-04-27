package main

import (
	"fmt"
	"go-rest-api/cmd/api/handlers"
	"go-rest-api/cmd/api/middlewares"
	"go-rest-api/cmd/api/validators"
	"go-rest-api/common"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	db, err := common.NewMySql()
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

	e.Validator = &validators.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middlewares.CustomMiddleware)

	app.routes()

	e.Logger.Fatal(e.Start(appAddress))
}
