package routes

import (
	middlewarecontroller "my-echo-app/controllers/middleware_controller"
	"my-echo-app/controllers/user_controller"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, store *sessions.CookieStore) {
	// 비보호된 라우트
	e.POST("/users", user_controller.CreateAccount)
	e.POST("/user/login", user_controller.Login)

	authGroup := e.Group("/user")
	authGroup.Use(middlewarecontroller.SessionAuth(store))

	authGroup.GET("/profile", user_controller.GetProfile)
}
