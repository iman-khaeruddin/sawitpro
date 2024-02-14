package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(id int) (string, error) {
	// Set custom claims
	claims := &AuthData{
		id,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}
