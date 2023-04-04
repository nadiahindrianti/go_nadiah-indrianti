package middlewares

import (
	"ORM_MVC/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
