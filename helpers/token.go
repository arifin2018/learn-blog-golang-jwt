package helpers

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var API_SECRET = Getenv("API_SECRET", "rahasiasekali")

func GenerateToken(user_id uint) (string, error) {
	token_lifespan, err := strconv.Atoi(Getenv("TOKEN_HOUR_LIFESPAN", "1"))

	if err != nil {
        return "", err
    }

	claims := jwt.MapClaims{}
	claims["authorized"] = true
    claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(API_SECRET))
}