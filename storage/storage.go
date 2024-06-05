package storage

import (
	"fitness-api/types"
	"fmt"
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
	query :=
		`INSERT INTO measurements 
			(user_id, weight, height, body_fat, created_at) 
			VALUES ($1, $2, $3, $4, $5) 
		RETURNING id`
	err := db.QueryRow(query, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err != nil {
		return measurement, err
	}
	return measurement, nil
}

// Updates the measurement in the database by user id and measurement id
func EditMeasurement(measurement types.Measurements) (types.Measurements, error) {
	db := GetDB()
	query :=
		`UPDATE measurements 
			SET 
				weight = $1, 
				height = $2, 
				body_fat = $3 
			WHERE 
				user_id = $4 
				AND 
				id = $5 
			RETURNING id`
	err := db.QueryRow(query, measurement.Weight, measurement.Height, measurement.BodyFat, measurement.UserId, measurement.Id).Scan(&measurement.Id)
	if err != nil {
		return measurement, err
	}
	return measurement, nil
}

// Deletes a measurement by its id and user id
func DeleteMeasurement(id, uid int) error {
	db := GetDB()
	query :=
		`DELETE FROM measurements
			WHERE 
				id = $1 
				AND 
				user_id = $2`
	result, err := db.Exec(query, id, uid)
	if err != nil {
		return err
	}

	// check if any rows were deleted
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsDeleted == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
