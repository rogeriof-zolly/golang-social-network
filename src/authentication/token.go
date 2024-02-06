package authentication

import (
	"devbook/src/config"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(6 * time.Hour).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, retrieveValidationKey)
	if err != nil {
		return err
	}

	fmt.Println(token)
	return nil
}

func extractToken(r *http.Request) string {
	tokenHeader := r.Header.Get("Authorization")

	splitToken := strings.Split(tokenHeader, " ")

	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}

func retrieveValidationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
