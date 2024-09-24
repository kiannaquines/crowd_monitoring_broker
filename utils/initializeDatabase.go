package utils

import (
	"database/sql"
	"log"
    "os"
    "fmt"
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

    _, dberr = db.Exec(`
        CREATE TABLE IF NOT EXISTS devices (
            device_id VARCHAR(36) PRIMARY KEY NOT NULL,
            device_addr TEXT NOT NULL,
            timestamp DATETIME NOT NULL,
            is_randomized BOOLEAN NOT NULL,
            device_power INTEGER NOT NULL,
            ssid TEXT NOT NULL,
            frame_type TEXT NOT NULL,
            zone TEXT NOT NULL,
            timestamp_created DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)

    if dberr != nil {
        log.Fatal(dberr)
    }

    log.Println("Database initialized")
}