package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte

func init() {
	secret = []byte(os.Getenv("JWT_SECRET"))
}

type JwtMapClaims struct {
	jwt.RegisteredClaims
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Exp      *jwt.NumericDate
}

func CreateToken(userId string, username string) (accessToken string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtMapClaims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   userId,
		},
		Exp: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	accessToken, err = jwtToken.SignedString(secret)
	return
}
