package routing

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("POST /t", TransactionHandler)
}
