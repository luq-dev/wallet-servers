package main

import (
	"log"
	"net/http"
	"servers/routing"
)

func main() {
	routing.RegisterRoutes()
	log.Println("Server is running on http://localhost:8080")
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Server Failed:%s\n", err)
	}
}
