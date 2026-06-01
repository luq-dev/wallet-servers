package routing

import (
	"auth/data"
	"auth/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Hello World

func hello_world(w http.ResponseWriter, req *http.Request) {

	t, err := services.GetToken(req.Header)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	tmap, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	uid, ok := tmap["uid"].(float64)
	if !ok {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	r := data.DB.QueryRow("SELECT fullname, email FROM users WHERE id = $1", int64(uid))

	var user data.User
	if err := r.Scan(&user.Name, &user.Email); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"name": user.Name, "email": user.Email})

}

// app routes

func addUser(w http.ResponseWriter, req *http.Request) {
	var user data.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Email != "" && user.Name != "" {
		_, err := data.DB.Exec("INSERT INTO users(fullname, email, password, phone_number) values ($1,$2,$3,$4)", user.Name, user.Email, string(password), user.PhoneNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Missing Data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(([]byte("user added")))
}

func getUser(w http.ResponseWriter, req *http.Request) {
	var user data.User

	w.Header().Set("Content Type", "application/json")

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	row := data.DB.QueryRow("SELECT fullname, email FROM users WHERE email=$1", user.Email)

	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		http.Error(w, "{\"Error\": \"User Not Found\"}", http.StatusNotFound)
		return
	}

	enc_err := json.NewEncoder(w).Encode(map[string]string{"email": user.Email, "fullname": user.Name})
	if enc_err != nil {
		panic(enc_err)
	}
}

func getAuthUser(w http.ResponseWriter, req *http.Request) {
	var user data.User
	var uid int64
	var p0 string

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	r := data.DB.QueryRow("SELECT id, password from users WHERE email = $1", user.Email)

	err := r.Scan(&uid, &p0)

	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(p0), []byte(user.Password)); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else {
		token, err := services.GenerateToken(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, token)
	}
}

func addAccount(w http.ResponseWriter, req *http.Request) {

	t, err := services.GetToken(req.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	uid, ok := mapClaims["uid"].(float64)
	if !ok {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	var acc data.Account
	dec_err := json.NewDecoder(req.Body).Decode(acc)

	if err != nil {
		http.Error(w, dec_err.Error(), http.StatusBadRequest)
		return
	}

	if acc.Name != "" && acc.Type != 0 {

		_, err := data.DB.Exec("INSERT INTO accounts(user_id, account_name, account_type) VALUES ($1, $2, $3)", int64(uid), acc.Name, acc.Type)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func getUserAccounts(w http.ResponseWriter, req *http.Request) {

}

func sendMoney(w http.ResponseWriter, req *http.Request) {

}

func fundWallet(w http.ResponseWriter, req *http.Request) {

}
