package main

import (
	"fitness-api/api"
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

	e.POST("/users", api.HandleCreateUser)
	e.PUT("/users/:id", api.HandleEditUser)

	e.POST("/measurements", api.HandleCreateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
