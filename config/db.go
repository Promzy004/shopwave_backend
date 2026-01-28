package config

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB () {
	DSN := GetDBDSN()

	DN := os.Getenv("DriverName")

	db, err := sql.Open(DN, DSN)
	if err != nil {
		panic(err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DB = db
}

func InitDB () {
	file, err := os.Open("schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	query, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	_ , err = DB.Exec(string(query))
	if err != nil {
		log.Fatal("Error creating table")
	}

}


func FreshMigrate() {
    file, err := os.Open("fresh.sql")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    queryBytes, err := io.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    _, err = DB.Exec(string(queryBytes))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database migrated fresh!")
}
