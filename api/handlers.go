package api

import (
	"fitness-api/storage"
	"fitness-api/types"
	"fitness-api/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func HandleCreateUser(c echo.Context) error {

	// bind request body to user struct
	user := types.User{}
	c.Bind(&user)

	// add new user in the database
	newUser, err := storage.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// generate jwt token to pass into the client
	token, err := utils.GenToken(newUser.Id, newUser.Name, newUser.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{"token": token})
}

func HandleCreateMeasurement(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*types.JwtCustomClaim)

	measurement := types.Measurements{}
	if err := c.Bind(&measurement); err != nil {
		return err
	}
	measurement.UserId = claims.UID

	// add measurement to database
	newMeasurement, err := storage.CreateNewMeasurement(measurement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newMeasurement)
}

// func HandleEditMeasurement(c echo.Context) error {

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	userToken := c.Get("user").(*jwt.Token)
// 	claims := userToken.Claims.(*types.JwtCustomClaim)

// 	measurement := types.Measurements{}
// 	if err := c.Bind(&measurement); err != nil {
// 		return err
// 	}
// 	measurement.UserId = claims.UID
// 	measurement.Id = id

// 	// TODO : edit measurement in DB
// 	return c.JSON(http.StatusOK, echo.Map{"message": measurement})
// }
