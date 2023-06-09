package controllers

import (
	"ORM_MVC/config"
	"ORM_MVC/lib/database"
	"ORM_MVC/models"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	users, err := database.GetUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"users":   users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user, err := database.CreateUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success created user",
		"data":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	deleteUser, err := database.DeleteUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"users":   deleteUser,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user, err := database.UpdateUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updated user",
		"data":    user,
	})
}

// method book

// get all book
func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all book",
		"data":    books,
	})
}

// get user by id
func GetBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	book, err := database.GetBook(bookID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"data":    book,
	})
}

// create new user
func CreateBookController(c echo.Context) error {
	book, err := database.CreateBook(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success created book ",
		"data":    book,
	})
}

// delete user by id
func DeleteBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	book, err := database.DeleteBook(bookID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
		"data":    book,
	})
}

// update user by id
func UpdateBookController(c echo.Context) error {
	book, err := database.UpdateBook(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id",
		"data":    book,
	})
}

func LoginUserControllers(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, err := database.LoginUsers(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := database.GetDetailUsers(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
