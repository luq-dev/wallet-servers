package routes

import (
	// auth_data "auth/data"
	// "auth/services"
	// "encoding/json"
	"net/http"
	// "transaction/data"

	// "github.com/golang-jwt/jwt/v5"
)

func sendMoney(w http.ResponseWriter, r *http.Request){
	// t, err := services.GetToken(r.Header)
	// if err != nil {
	// 	http.Error(w, "Unauthorized User", http.StatusUnauthorized)
	// 	return
	// }

	// tmap, ok := t.Claims.(jwt.MapClaims)
	// if !ok {
	// 	http.Error(w, "Invalid User Claims", http.StatusUnauthorized)
	// 	return
	// }

	// var t data.Transaction

	// dec_err := json.NewDecoder(r.Body).Decode(t)
	// if dec_err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	

}

func fundWallet(w http.ResponseWriter, r *http.Request){

}