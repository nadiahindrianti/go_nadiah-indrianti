package controllers

import (
	"Task_2/configs"
	"Task_2/helpers"
	database "Task_2/lib/database"
	"Task_2/middlewares"
	"Task_2/models"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := configs.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Proses Register Gagal",
			Message: "Terdeteksi Eror",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Proses Registrasi Berhasil",
		Message: "Successfuly Registrasi User",
	})

}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, err := database.LoginUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reflectValue := reflect.ValueOf(result)
	userID := reflectValue.FieldByName("ID").Interface().(uint)
	userName := reflectValue.FieldByName("Name").Interface().(string)
	userEmail := reflectValue.FieldByName("Email").Interface().(string)
	userContact := reflectValue.FieldByName("Contact").Interface().(string)
	userAlamat := reflectValue.FieldByName("Alamat").Interface().(string)
	userRole := reflectValue.FieldByName("Role").Interface().(string)

	token, err := middlewares.CreateTokenUser(int(userID), userName, userEmail, userContact, userAlamat, userRole)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userResponse := models.UserResponse{ID: userID, Name: userName, Email: userEmail, Role: userRole, Token: token}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Proses Login Berhasil",
		Message: "Successfuly Login User",
		Data:    userResponse,
	})
}
