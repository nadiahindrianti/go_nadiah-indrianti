package middlewares

import (
	"ORM_MVC/config"
	"ORM_MVC/models"

	"github.com/labstack/echo"
)

var db = config.DB

func BasicAutoDB(username, password string, c echo.Context) (bool, error) {
	var user models.User
	err := db.Where("id = ? AND password = ?", user.ID, user.Password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
