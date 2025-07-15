package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}
