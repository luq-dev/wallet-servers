package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"storage/database"
	"user/data"
	"user/services/auth"
	"user/services/dao"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var userDAO = dao.NewUserDAO(database.DB)

func signup(w http.ResponseWriter, req *http.Request) {
	var user data.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := userDAO.AddUser(&user)

	if err == nil {
		tk, e := auth.GenerateToken(id)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}

		res, enc_err := json.Marshal(map[string]string{"token": tk})
		if enc_err != nil {
			http.Error(w, enc_err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	var user data.User
	var uid int64
	var p0 string

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := database.DB.QueryRow("SELECT id, password from users WHERE email = $1", user.Email).Scan(&uid, &p0)

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

func pingUser(w http.ResponseWriter, req *http.Request) {

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

	user, err := userDAO.GetUserByID(int64(uid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	enc_err := json.NewEncoder(w).Encode(user)
	if err != nil{
		http.Error(w, enc_err.Error(), http.StatusInternalServerError)
	}
}
