package network

import (
	"encoding/json"
	"log"
	"net/http"
	"message/models"
	"user/services/auth"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)



var u = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	var messageBuffer models.Message
	t, t_err := auth.GetToken(r.Header)

	if t_err != nil {
		http.Error(w, t_err.Error(), http.StatusUnauthorized)
		return
	}

	_, ok := t.Claims.(jwt.MapClaims) // TODO: claims to be made

	if !ok {
		http.Error(w, "Invalid Tokens", http.StatusUnauthorized)
		return
	}

	conn, err := u.Upgrade(w, r, nil) // TODO:
	if err != nil {
		log.Println(err.Error())
		return
	}

	go func() { // Recieve Listener
		for {
			_, msg, err := conn.ReadMessage()

			if err != nil {
				log.Println(err.Error())
				continue
			}

			dec_err := json.Unmarshal(msg, &messageBuffer)
			if dec_err != nil {
				log.Println(dec_err.Error())
				continue
			}
		}
	}()

	go func() { // Sending Channel
		for {
			message := "Test"
			if err := conn.WriteMessage(websocket.BinaryMessage, []byte(message)); err != nil {
				log.Printf("Send Error: %v", err.Error())
			}
		}
	}()
}
