package routes

import (
	"my-echo-app/controllers/blog_controller"
	middlewarecontroller "my-echo-app/controllers/middleware_controller"
	"my-echo-app/controllers/user_controller"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, store *sessions.CookieStore) {

	e.POST("/user", user_controller.CreateAccount)
	e.POST("/user/login", user_controller.Login)
	e.POST("/user/logout", user_controller.Logout)

	// protected routes
	authGroup := e.Group("/user")
	authGroup.Use(middlewarecontroller.SessionAuth(store))

	authGroup.GET("/profile", user_controller.GetProfile)

	blogGroup := e.Group("/blog")
	blogGroup.Use(middlewarecontroller.SessionAuth(store))

	blogGroup.POST("", blog_controller.CreateBlog)
	blogGroup.GET("", blog_controller.GetBlogs)
	blogGroup.DELETE("/:id", blog_controller.DeleteBlog)
	blogGroup.PUT("/:id", blog_controller.UpdateBlog)
}
