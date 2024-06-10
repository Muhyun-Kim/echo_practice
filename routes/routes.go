package routes

import (
	"my-echo-app/controllers/root_controller"
	"my-echo-app/controllers/user_controller"
	"os"

	jwtMiddleware "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", root_controller.GetRoot)
	e.POST("/user/create-account", user_controller.CreateAccount)
	e.POST("/user/login", user_controller.Login)

	r := e.Group("/user")
	r.Use(jwtMiddleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.GET("/profile", user_controller.GetProfile)
}
