package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv () {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env file")
	}
}

func GetEnv (key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}


func GetDBDSN () string {

	user := GetEnv("DB_Username", "root")
	host := GetEnv("DB_Host", "127.0.0.1")
	port := GetEnv("DB_Port", "3306")
	dbName := GetEnv("DB_DATABASE", "")
	password := GetEnv("DB_PASSWORD", "")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", user, password, host, port, dbName)
}
