package main

import (
	"log"
	"net/http"
	
	"user/database"
	. "transaction/routing"
	. "transaction/services"
)

func main(){
	t := NewTrasactionProcessor(database.DB)
	go t.FetchAndProcessTransactions()

	RegisterRoutes()

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Printf("Server Failed: %s\n", err)
	}
}