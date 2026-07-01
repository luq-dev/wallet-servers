package main

import (
	"log"
	"net/http"
	
	. "transaction/routing"
)

func main(){

	RegisterRoutes()

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Printf("Server Failed: %s\n", err)
	}
}