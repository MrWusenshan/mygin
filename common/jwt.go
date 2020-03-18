package common

import (
	"github.com/dgrijalva/jwt-go"
	"mygin/modle"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user modle.User) (string, error) {
	expirationTime := time.Now().Add(7 * time.Hour * 24).Unix()
	claims := Claims{UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "MrWusenshan",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//解析token函数
func ParesToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
