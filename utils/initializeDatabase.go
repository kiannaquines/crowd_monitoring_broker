package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func InitializeDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var dberr error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, dberr = sql.Open("mysql", dsn)

	if dberr != nil {
		log.Fatal(dberr)
	}

	if dberr := db.Ping(); dberr != nil {
		log.Fatal(dberr)
	}

	if dberr != nil {
		log.Fatal(dberr)
	}

	log.Println("Database initialized")
}
