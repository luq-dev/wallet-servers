package data

import (
	"os"
	"log"
	"time"
	"strconv"
	"database/sql"
	"github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	cfg := pq.Config{
		Database:       getEnv("DB_NAME", "green_test"),
		Host:           getEnv("DB_HOST", "localhost"),
		Port:           uint16(getEnvInt("DB_PORT", 5432)),
		User:           getEnv("DB_USER", "dl"),
		Password:       getEnv("DB_PASSWORD", ""),
		ConnectTimeout: time.Duration(getEnvInt("DB_TIMEOUT", 5)) * time.Second,
	}

	connector, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	DB := sql.OpenDB(connector)

	return DB
}

var DB = ConnectDB()

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer for %s: %v, using default %d\n", key, err, defaultValue)
		return defaultValue
	}
	return intValue
}
