package auth

import "github.com/labstack/echo/v4"

// AuthRouter :
type AuthRouter struct{}

// Init : Init Group
func (ctrl AuthRouter) Init(g *echo.Group) {
	g.POST("/register", ctrl.Register)
	g.POST("/login", ctrl.Login)
}
