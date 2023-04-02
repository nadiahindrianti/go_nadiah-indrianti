package routes

import (
	"ORM_MVC/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controllers.GetUserController)
	e.GET("/user/:id", controllers.GetUsersController)
	e.POST("users", controllers.CreateUserController)
	e.DELETE("users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("books", controllers.CreateBookController)
	e.DELETE("books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	return e
}
