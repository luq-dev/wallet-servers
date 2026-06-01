package main

import (
	"auth/routing"
	"log"
	"net/http"
)

func main() {
	routing.RegisterRoutes()
	log.Println("Server is running on http://0.0.0.0:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Server Failed:%s\n", err)
	}
}
