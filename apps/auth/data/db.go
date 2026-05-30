package data

import (
	"database/sql"
	"log"
	"time"

	// "github.com/golang-migrate/migrate/v4"
	"github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)


func ConnectDB() *sql.DB{

	cfg := pq.Config{
		Database:       "green_test",
		Host:           "localhost",
		Port:           5432,
		User:           "dl",
		ConnectTimeout: 5 * time.Second,
		Password:       "cathereen",
	}

	connector, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	DB := sql.OpenDB(connector)

	return DB
}

var DB = ConnectDB()

var AccountTypes map[string]int64 = map[string]int64{
	"TEST": 1,
}
