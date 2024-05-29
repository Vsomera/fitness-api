package handlers

import (
	"fitness-api/storage"
	"fitness-api/types"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := types.User{}
	c.Bind(&user)
	newUser, err := storage.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}

func EditUser(c echo.Context) error {
	// get param from request url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// bind request body and current time to user struct
	user := types.User{}
	c.Bind(&user)
	user.Id = id
	user.UpdatedAt = time.Now()

	// edit user
	err = storage.EditUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "User updated")
}
