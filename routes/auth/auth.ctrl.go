package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//Login : Login Router
func (AuthRouter) Login(c echo.Context) error {
	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	if body.Username == "user" && body.Password == "1234" {
		return c.HTML(http.StatusOK, "<h1>Hello user!</h1>")
	}

	return c.HTML(http.StatusBadRequest, "<h1>ERROR</h1>")
}
