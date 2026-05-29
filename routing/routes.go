package routing

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("GET /u", hello_world)
	http.HandleFunc("POST /u/add", addUser)
	http.HandleFunc("POST /u/get", getUser)
	http.HandleFunc("POST /u/signin", getAuthUser)
	http.HandleFunc("POST /acc/add", addAccount)		// to be worked on
	http.HandleFunc("POST /acc/get", getUserAccounts)	// just accounts associated with the use
	http.HandleFunc("POST /t/send", sendMoney)
	http.HandleFunc("POST /t/fund", fundWallet)
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
