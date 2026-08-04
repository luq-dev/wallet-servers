package main

import (
	. "carding/routing"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	RegisterRoutes()

	fmt.Println("Server Listening at 127.0.0.1:8083")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Println("Server Error", err.Error())
	}
}

// Mock function that makes a pool of 100 card numbers for use
// later to avoid slow DB transactions (i think its stupid and
// might change later) so that all one has to do is assign a
// card to a user rather than iterate everytime. This saves
// from using Queues.
func make100Cards(db *sql.DB, network int) error{
	// in normal cases cards will be made from a specified
	// network therefore requiring a requesting procedure but
	// that is a story for another time
	_, err :=db.Exec("SELECT make100cards($1)", network)
	return err
}
