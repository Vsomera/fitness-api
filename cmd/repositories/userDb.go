package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
)

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}
