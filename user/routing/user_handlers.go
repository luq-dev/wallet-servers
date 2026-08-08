package routing

import (
	"encoding/json"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	id, err := userDAO.AddUser(&user)

	if err == nil {
		tk, e := auth.GenerateToken(id)
		if e != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error":e.Error()})
			return
		}

		res, enc_err := json.Marshal(map[string]string{"token": tk})
		if enc_err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error":enc_err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error":err.Error()})
		return
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	var user data.User
	var uid int64
	var p0 string

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error":err.Error()})
		return
	}

	err := database.DB.QueryRow("SELECT id, password from users WHERE email = $1", user.Email).Scan(&uid, &p0)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error":"User Not Found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(p0), []byte(user.Password)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid password"})
		return
	} else {
		token, err := auth.GenerateToken(uid)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func pingUser(w http.ResponseWriter, req *http.Request) {

	t, err := auth.GetToken(req.Header)
	if err != nil {
		w.Header().Set("Content-Type", "applicatio/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid Token"})
		return
	}

	tmap, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Token Claims"})
		return
	}

	uid, ok := tmap["uid"].(float64)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid Token Claim"})
		return
	}

	user, err := userDAO.GetUserByID(int64(uid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
