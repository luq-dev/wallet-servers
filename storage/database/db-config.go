package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	Db, err := sql.Open(
		"postgres",
		"postgres://user:password@localhost:5432/mydb?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	DB = Db
}
