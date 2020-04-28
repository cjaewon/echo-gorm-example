package main

import (
	"github.com/cjaewon/echo-gorm-example/database"
	"github.com/cjaewon/echo-gorm-example/lib"
	"github.com/cjaewon/echo-gorm-example/lib/middlewares"
	"github.com/cjaewon/echo-gorm-example/routes"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	godotenv.Load(".env.development")
}

func main() {
	db := database.Connect()
	defer db.Close()

	e := echo.New()

	e.Validator = &lib.CustomValidator{Validator: validator.New()}

	e.Use(middlewares.ContextDB(db))

	routes.Routes(e.Group(""))

	e.Logger.Fatal(e.Start(":3000"))
}
