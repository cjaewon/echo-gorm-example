package main

import (
	"github.com/cjaewon/echo-gorm-example/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Routes(e.Group(""))

	e.Logger.Fatal(e.Start(":3000"))
}
