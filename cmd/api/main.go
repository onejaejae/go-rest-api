package main

import (
	"fmt"
	"go-rest-api/cmd/api/handlers"
	"go-rest-api/cmd/api/rotues"
	"go-rest-api/cmd/api/services"
	"go-rest-api/cmd/api/validators"
	"go-rest-api/common"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	service := services.NewService(db)
	handler := handlers.NewHandler(service)

	e.Validator = &validators.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())

	rotues.RegisterAppRoutes(e, handler.AppHandler)
	rotues.RegisterAuthRoutes(e, handler.AuthHandler)

	e.Logger.Fatal(e.Start(appAddress))
}
