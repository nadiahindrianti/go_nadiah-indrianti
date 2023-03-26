package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User
users.Id = "1234567"
users.Email = "rasyaaruq123@gmail.com"
user.Nama = "Rasya Aruq"
user.Password = "123rasya"


// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	user := new(Users)
	user.Id = c.QuerryParam("keywords")
	user.Email = c.QuerryParam("keyword")
	user.Nama = "Jamali Rasyid"
	user.Password = "abc123"
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"user":     user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	user := new(Users)
	user.Id = c.Param("id")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"user":     user,
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	user := new(Users)
	c.Bind(user)

	user.Id = c.Param("id")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"user":     user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUsersController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
