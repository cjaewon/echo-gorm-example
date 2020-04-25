package main

import (
	"github.com/cjaewon/echo-gorm-example/database"
	"github.com/cjaewon/echo-gorm-example/lib/middleware"
	"github.com/cjaewon/echo-gorm-example/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env.development")

	db := database.Connect()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.ContextDB(db))

	routes.Routes(e.Group(""))

	e.Logger.Fatal(e.Start(":3000"))

}
