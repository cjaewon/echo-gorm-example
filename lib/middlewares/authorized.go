package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Authoriszed : Check Auth
func Authoriszed(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenCookie, err := c.Cookie("token")
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		c.Set("tokenCookie", tokenCookie)
		return next(c)
	}
}
