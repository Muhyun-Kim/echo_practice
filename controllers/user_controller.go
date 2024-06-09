package controllers

import (
	"my-echo-app/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{
        "data": "hello",
    })
}

func Ping(c echo.Context) error {
    sqlDB, err := database.DB.DB()
    if err != nil {
        return c.String(http.StatusInternalServerError, "MySQL 연결 실패")
    }
    if err := sqlDB.Ping(); err != nil {
        return c.String(http.StatusInternalServerError, "MySQL 연결 실패")
    }
    return c.String(http.StatusOK, "MySQL 연결 성공")
}
