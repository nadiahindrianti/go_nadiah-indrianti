package controllers

import (
	"Task_1/configs"
	"Task_1/helpers"
	database "Task_1/lib/database"
	"Task_1/models"
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

func GetItemById(id any) (models.Item, error) {
	var item models.Item

	err := configs.DB.Where("id = ?", id).First(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func GetItemByCategori(name string) (models.Item, error) {
	var item models.Item

	err := configs.DB.Where("name = ?", name).First(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func GetItemController(c echo.Context) error {
	Category, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item, err := database.GetItembyCategory(Category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Get Item",
		"products": item,
	})
}

func GetItemControllerbyName(c echo.Context) error {
	Name, err := strconv.Atoi(c.Param("name"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item, err := database.GetItembyName(Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Success Get Item",
		"products": item,
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
