package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
)

type AuthData struct {
	ID         int  `json:"id"`
	Authorized bool `json:"authorized"`
	jwt.RegisteredClaims
}

func (authData *AuthData) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, authData)
	}
	return err
}

func GetData(c echo.Context) (AuthData, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return AuthData{}, errors.New("JWT token missing or invalid")
	}

	var userClaim AuthData
	jwtToken, err := jwt.ParseWithClaims(token.Raw, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// Checking token validity
	if !jwtToken.Valid {
		fmt.Println("invalid token")
	}

	c.Set("authData", userClaim)
	fmt.Printf("Parsed User Claim: %d\n", userClaim.ID)

	return userClaim, nil
}
