package routes

import (
	"ORM_MVC/controllers"
	m "ORM_MVC/middlewares"

	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	m.LogMiddlewares(e)
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

	eAuth := e.Group("/auth")
	eAuth.Use(echoMid.BasicAuth(m.BasicAutoDB))
	eAuth.GET("/users", controllers.GetUserController)
	eAuth.GET("/user/:id", controllers.GetUsersController)
	eAuth.POST("users", controllers.CreateUserController)
	eAuth.POST("users/login", controllers.CreateUserController)
	eAuth.DELETE("users/:id", controllers.DeleteUserController)
	eAuth.PUT("/users/:id", controllers.UpdateUserController)
	eAuth.GET("/books", controllers.GetBooksController)
	eAuth.GET("/books/:id", controllers.GetBookController)
	eAuth.POST("books", controllers.CreateBookController)
	eAuth.DELETE("books/:id", controllers.DeleteBookController)
	eAuth.PUT("/books/:id", controllers.UpdateBookController)

	eJwt := e.Group("/jwt")
	eJwt.GET("/users", controllers.GetUserDetailController)
	eJwt.GET("/user/:id", controllers.GetUsersController)
	eJwt.POST("users", controllers.CreateUserController)
	eJwt.DELETE("users/:id", controllers.DeleteUserController)
	eJwt.PUT("/users/:id", controllers.UpdateUserController)
	eJwt.GET("/books", controllers.GetBooksController)
	eJwt.GET("/books/:id", controllers.GetBookController)
	eJwt.POST("books", controllers.CreateBookController)
	eJwt.DELETE("books/:id", controllers.DeleteBookController)
	eJwt.PUT("/books/:id", controllers.UpdateBookController)
	return e
}
