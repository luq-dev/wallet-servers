package services

import (
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("1234567345678")

func GenerateToken(uid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(SECRET)
}

func VerifyToken(t string) (*jwt.Token, error){
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return SECRET, nil
	})
}
