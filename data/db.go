package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// "github.com/golang-migrate/migrate/v4"
	"github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

var Db *sql.DB

// func runMigrations() error {
// 	m, err := migrate.New(
// 		"file://migrations",
// 		"postgres://dl:cathereen@localhost:5432/green_test?sslmode=disable",
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer m.Close()

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		return err
// 	}

// 	return nil
// }

func init() {

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

	Db := sql.OpenDB(connector)

	err = Db.Ping()
	if err == nil {
		fmt.Println("Database Connected Successfully")
	}
}

// DAO - Data Access Objects
