package handlers

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMeasurement(c echo.Context) error {
	var measurement models.Measurements
	if err := c.Bind(&measurement); err != nil {
		return err
	}
	newMeasurement, err := repositories.CreateNewMeasurement(measurement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newMeasurement)
}
