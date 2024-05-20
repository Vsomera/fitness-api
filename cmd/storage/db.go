package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	psqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	// establish database connection
	var err error
	db, err = sql.Open("postgres", psqlDbInfo)
	if err != nil {
		panic(err)
	}

	// check if database connected successfully
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}
