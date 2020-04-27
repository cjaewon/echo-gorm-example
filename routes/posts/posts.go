package posts

import (
	"github.com/cjaewon/echo-gorm-example/lib/middlewares"
	"github.com/labstack/echo/v4"
)

// PostsRouter : PostsRouter struct
type PostsRouter struct{}

// Init : Init Router
func (ctrl PostsRouter) Init(g *echo.Group) {
	g.GET("/read/:id", ctrl.Read)
	g.POST("/create", ctrl.Create, middlewares.Authoriszed)
	g.DELETE("/delete/:id", ctrl.Delete, middlewares.Authoriszed)
}
