package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	storage.InitDB()

	e := echo.New()
	e.GET("/", handlers.Home)

	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.EditUser)

	e.POST("/measurements", handlers.CreateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
