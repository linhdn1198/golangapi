package util

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var hmacSampleSecret []byte

func Generate(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(hmacSampleSecret)

	return tokenString
}

func Parse(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}