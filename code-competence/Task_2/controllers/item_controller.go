package controllers

import (
	"Task_2/configs"
	"Task_2/helpers"
	database "Task_2/lib/database"
	"Task_2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateItemController(c echo.Context) error {
	item := models.Item{}
	c.Bind(&item)

	if err := configs.DB.Save(&item).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Proses Create Item Gagal",
			Message: "Terdeteksi Eror",
		})
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Proses Create Berhasil",
		Message: "Successfuly Create Product",
	})

}

func GetItemControllerAll(c echo.Context) error {
	items, err := database.GetItemController()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Get All Items",
		"products": items,
	})
}

func GetItemControllerbyId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var item models.Item
	if err = configs.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Item by ID",
		Message: "Successfuly",
		Data:    item,
	})
}

func GetItemControllerbyCategoryId(c echo.Context) error {
	categoryid, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var item models.Item
	if err = configs.DB.Where("categoryid = ?", categoryid).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Item by CategoryID",
		Message: "Successfuly",
		Data:    item,
	})
}

func GetItemControllerbyName(c echo.Context) error {
	name, err := strconv.Atoi(c.Param("name"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var item models.Item
	if err = configs.DB.Where("name = ?", name).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseData{
		Status:  "Success Get Item by CategoryID",
		Message: "Successfuly",
		Data:    item,
	})
}

func UpdateItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}
	var item models.Item
	if err := configs.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	c.Bind(&item)
	if err := configs.DB.Model(&item).Updates(item).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.ResponseNotData{
		Status:  "Success Upload Data Item",
		Message: "Successfuly",
	})

}

func DeleteItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	var item models.Item
	if err := configs.DB.First(&item, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
			Status:  "Error",
			Message: "Data Tidak Tersedia",
		})
	}

	if err := configs.DB.Delete(&item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal Delete Item Data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusBadRequest, helpers.ResponseNotData{
		Status:  "Succes Delete Data Item",
		Message: "Succesfuly",
	})
}
