package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"finance/models"
	"user/data"
	"user/database"
	"user/services/auth"
	"user/services/dao"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var userDAO = dao.NewUserDAO(database.DB)

// Hello World

func hello_world(w http.ResponseWriter, req *http.Request) {

	t, err := auth.GetToken(req.Header)
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

	r := database.DB.QueryRow("SELECT fullname, email FROM users WHERE id = $1", int64(uid))

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

	if err := userDAO.AddUser(&user); err == nil {
		w.WriteHeader(http.StatusCreated)
		w.Write(([]byte("user added")))
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func getUser(w http.ResponseWriter, req *http.Request) {
	var user data.User

	w.Header().Set("Content Type", "application/json")

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := userDAO.GetUserByEmail(user.Email)
	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
	}

	enc_err := json.NewEncoder(w).Encode(map[string]string{"email": u.Email, "fullname": u.Name})
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

	r := database.DB.QueryRow("SELECT id, password from users WHERE email = $1", user.Email)

	err := r.Scan(&uid, &p0)

	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(p0), []byte(user.Password)); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else {
		token, err := auth.GenerateToken(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, token)
	}
}

func addAccount(w http.ResponseWriter, req *http.Request) {

	t, err := auth.GetToken(req.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid Claims", http.StatusUnauthorized)
		return
	}

	uid, ok := mapClaims["uid"].(float64)
	if !ok {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	var acc models.Account
	dec_err := json.NewDecoder(req.Body).Decode(&acc)

	if dec_err != nil {
		http.Error(w, dec_err.Error(), http.StatusBadRequest)
		return
	}

	if acc.Name != "" && acc.Type != "" {

		_, err := database.DB.Exec("INSERT INTO accounts(user_id, account_name, account_type) VALUES ($1, $2, $3)", int64(uid), acc.Name, acc.Type)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func getUserAccounts(w http.ResponseWriter, req *http.Request) {
	t, err := auth.GetToken(req.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid Map Claims", http.StatusUnauthorized)
		return
	}

	uid, ok := mapClaims["uid"].(float64)
	if !ok {
		http.Error(w, "Invalid Map claims", http.StatusUnauthorized)
		return
	}

	rows, err := database.DB.Query("SELECT account_id, account_name, account_type FROM user_account_details WHERE user_id = $1", int64(uid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var accounts []models.Account

	for rows.Next() {
		var r1 int64
		var r2 string
		var r3 string
		rows.Scan(&r1, &r2, &r3)
		accounts = append(accounts, models.Account{ID: r1, Type: r2, Name: r3})
	}
	rows.Close()

	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
