package network

import (
	"encoding/json"
	"net/http"
	. "message/models"
	. "message/core"
	"user/database"
)

var ms = NewMessageService(database.DB)


func collect(w http.ResponseWriter, r *http.Request){

	var msg Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		http.Error(w, "Invalid Message", http.StatusBadRequest)
		return
	}
}