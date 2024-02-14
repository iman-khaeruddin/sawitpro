package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
	"sawitpro/handler"
	"sawitpro/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	e := echo.New()
	db := util.GormPostgres(os.Getenv("POSTGRES_DSN"))
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	// login
	loginHandler := handler.NewLoginHandler(db)
	loginHandler.LoginHandler(e)

	// register
	registerHandler := handler.NewRegisterHandler(db)
	registerHandler.RegisterHandler(e)

	// profile
	getProfileHandler := handler.NewProfileHandler(db)
	getProfileHandler.ProfileHandler(e)

	e.Logger.Fatal(e.Start(":8080"))
}
