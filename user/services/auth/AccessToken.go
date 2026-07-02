package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("1234567345678")

func GenerateToken(uid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(secret)
}

func GetToken(h http.Header) (*jwt.Token, error) {
	authString := h.Get("Authorization")
	tokenString := strings.TrimPrefix(authString, "Bearer ") // use [:7 slice in case of performance issues]

	if tokenString == authString {
		return nil, fmt.Errorf("Missing or invalid Authotrization Header")
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
}
