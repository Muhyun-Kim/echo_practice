package root_controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRoot(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"data": "home",
	})
}