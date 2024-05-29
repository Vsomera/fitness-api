package handlers

import (
	"fitness-api/storage"
	"fitness-api/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMeasurement(c echo.Context) error {
	var measurement types.Measurements
	if err := c.Bind(&measurement); err != nil {
		return err
	}
	newMeasurement, err := storage.CreateNewMeasurement(measurement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newMeasurement)
}
