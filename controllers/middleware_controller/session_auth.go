package middlewarecontroller

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func SessionAuth(store *sessions.CookieStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, err := store.Get(c.Request(), "session-name")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Failed to get session",
				})
			}

			email, ok := session.Values["email"].(string)
			if !ok || email == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Missing or invalid session",
				})
			}

			c.Set("email", email)

			return next(c)
		}
	}
}
