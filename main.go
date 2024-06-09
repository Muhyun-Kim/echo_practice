package main

import (
	"my-echo-app/database"
	"my-echo-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // 데이터베이스 연결
    database.Connect()

    // 라우트 설정
    routes.RegisterRoutes(e)

    // 서버 시작
    e.Logger.Fatal(e.Start(":8080"))
}
