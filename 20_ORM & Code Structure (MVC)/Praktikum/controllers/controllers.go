package controllers

import (
	"ORM_MVC/config"
	"ORM_MVC/models"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var users models.User

	if err := config.DB.First(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data user",
		"user":    users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var users models.User

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var users models.User

	if err := config.DB.Update(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if config.DB.Where("id = ?", id).Update(&users).RowsAffected == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "can't update data user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// method book

// get all book
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"users":   books,
	})
}

// get user by id
func GetBookController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var books models.Book

	if err := config.DB.First(&books, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data book",
		"user":    books,
	})
}

// create new user
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    book,
	})
}

// delete user by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var books models.Book

	if err := config.DB.Delete(&books, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update user by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	penerbit := c.Param("penerbit")
	var books models.Book

	if err := config.DB.Update(&books, id, penerbit).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if config.DB.Where("id = ?", id).Update(&books, id, penerbit).RowsAffected == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "can't update data book",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
	})
}
