package database

import (
	"ORM_MVC/config"
	"ORM_MVC/middlewares"
	"ORM_MVC/models"

	"github.com/labstack/echo"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(userID int) (interface{}, error) {
	var user models.User

	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	var user models.User
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(c echo.Context) (interface{}, error) {
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return nil, err
	}
	if err := c.Bind(&user); err != nil {
		return nil, err
	}
	config.DB.Save(&user)

	return user, nil
}

func DeleteUser(userID int) (interface{}, error) {
	var user models.User

	if err := config.DB.Where("id = ?", userID).Delete(&user).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("id = ? AND password = ?", user.ID, user.Password).First(&user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
