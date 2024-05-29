package main

import (
	"fitness-api/handlers"
	"fitness-api/storage"
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

	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.EditUser)

	e.POST("/measurements", handlers.CreateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
