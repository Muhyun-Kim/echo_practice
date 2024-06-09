package routes

import (
	"my-echo-app/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
    e.GET("/", controllers.Root)
    e.GET("/ping", controllers.Ping)
}
