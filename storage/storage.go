package storage

import (
	"fitness-api/types"
	"time"
)

// CreateUser creates a new user in the db with the given name, email, and password
func CreateUser(user types.User) (types.User, error) {
	db := GetDB()
	sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateNewMeasurement adds a new row into the db with the associated user_id
func CreateNewMeasurement(measurement types.Measurements) (types.Measurements, error) {
	db := GetDB()
	query := "INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := db.QueryRow(query, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err != nil {
		return measurement, err
	}
	return measurement, nil
}
