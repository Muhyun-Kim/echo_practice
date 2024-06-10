package main

import (
	"log"
	"my-echo-app/database"
	"my-echo-app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.Connect()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
