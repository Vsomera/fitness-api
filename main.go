package main

import (
	"fitness-api/api"
	"fitness-api/storage"
	"fitness-api/types"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	storage.InitDB()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", api.HandleCreateUser) // create a new user

	// RESTRICTED ROUTES (require bearer token to access)
	r := e.Group("/measurements")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.JwtCustomClaim)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
	r.Use(echojwt.WithConfig(config))

	r.POST("", api.HandleCreateMeasurement)
	// r.PUT("/:id", api.HandleEditMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
