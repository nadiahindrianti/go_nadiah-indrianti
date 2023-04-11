package routes

import (
	"clean_architecture/repository"

	"github.com/labstack/echo"
	echomid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	jwtAuth := e.Group("")
	jwtAuth.Use(echomid.JWT([]byte(repository.SECRET_JWT)))
	jwtAuth.GET("/users", controller.GetUserController)
	jwtAuth.Create("/users", controller.CreateUserController)
	return e
}
