package routing

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("GET /test", conn_test)
	http.HandleFunc("GET /u", hello_world)
	http.HandleFunc("POST /u/add", addUser)
	http.HandleFunc("POST /u/get", getUser)
	http.HandleFunc("POST /u/signin", getAuthUser)
	http.HandleFunc("POST /acc/add", addAccount)		// to be worked on
	http.HandleFunc("POST /acc/get", getUserAccounts)	// just accounts associated with the use
}

func conn_test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
/*

Get user details, 			[x]
get user auth (sign In)		[x]
add user					[x]

add Account					[]
get user accounts,			[]
get account details			[]

send money					[]
add money					[]
request money 				[]

*/
