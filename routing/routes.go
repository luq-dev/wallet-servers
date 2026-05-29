package routing

import "net/http"

func RegisterRoutes(){
	http.HandleFunc("GET /u", hello_world)
	http.HandleFunc("POST /addUser", addUser)
}

/*
Get user details,
get user accounts,
get account details
add Account
get user auth (sign In)
add user

*/

