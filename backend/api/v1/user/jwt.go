package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func generateJWT(id uuid.UUID, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
		"id": id,
        "username": username, 
        "exp": time.Now().Add(time.Hour * 72).Unix(), 
        })
	
	tokenString, err := token.SignedString(secretKey)
		if err != nil {
		return "", err
	 }

	return tokenString, nil
}