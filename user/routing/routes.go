package routing

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("GET /u", pingUser)
	http.HandleFunc("POST /u/signup", signup)
	http.HandleFunc("POST /u/signin", login)
	http.HandleFunc("POST /acc/add", addAccount)
	http.HandleFunc("POST /acc/get", getUserAccounts)
}