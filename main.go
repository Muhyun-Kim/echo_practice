package main

import (
	"log"
	middlewarecontroller "my-echo-app/controllers/middleware_controller"
	"my-echo-app/database"
	"my-echo-app/routes"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var store *sessions.CookieStore

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 환경 변수에서 해시 키와 암호화 키 가져오기
	hashKey := []byte(os.Getenv("SESSION_HASH_KEY"))
	blockKey := []byte(os.Getenv("SESSION_BLOCK_KEY"))

	store = sessions.NewCookieStore(hashKey, blockKey)
}

func main() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.Connect()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, err := store.Get(c.Request(), "session-name")
			if err != nil {
				c.Logger().Error("Session get error: ", err)
				return err
			}
			c.Set("session", session)
			return next(c)
		}
	})

	e.Use(middlewarecontroller.SessionAuth(store))

	routes.RegisterRoutes(e, store)

	e.Logger.Fatal(e.Start(":8080"))
}
