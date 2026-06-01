package routes

import "net/http"

func RegisterRoutes(){
	http.HandleFunc("POST /t/send", sendMoney)
	http.HandleFunc("POST /t/fund", fundWallet)
}
	