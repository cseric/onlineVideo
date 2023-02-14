package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	Id uint 	`json:"id"`
	jwt.RegisteredClaims
}

const TokenPrefix = "Bearer "

var secret = []byte(viper.GetString("app.token_secret"))

func ReleaseToken(id uint) (string, error) {
	expiresTime := time.Now().Add(time.Hour * 2)
	claims := Claims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "eric",
			Subject: "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	return token, claims, err
}

