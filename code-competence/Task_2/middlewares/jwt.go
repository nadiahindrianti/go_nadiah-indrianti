package middlewares

import (
	"Task_2/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateTokenUser(ID int, name string, email string, contact string, alamat string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = ID
	claims["name"] = name
	claims["email"] = email
	claims["contact"] = contact
	claims["alamat"] = alamat
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(int)
		return userId
	}
	return 0
}
