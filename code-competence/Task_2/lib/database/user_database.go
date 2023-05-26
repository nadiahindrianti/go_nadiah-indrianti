package database

import (
	"Task_2/configs"
	"Task_2/models"
)

func LoginUser(user models.User) (interface{}, error) {
	var err error

	if err = configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func RegisterUser(user models.User) (interface{}, error) {
	err := configs.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
