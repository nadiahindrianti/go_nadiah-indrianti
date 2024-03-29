package routes

import (
	"Task_2/controllers"
	m "Task_2/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	eUser := e.Group("/user")
	eUser.POST("/register", controllers.RegisterUserController)
	eUser.POST("/login", controllers.LoginUserController)

	eItem := e.Group("/item")
	eItem.GET("", controllers.GetItemControllerAll)
	eItem.GET("/:id", controllers.GetItemControllerbyId)
	eItem.GET("/category/:id", controllers.GetItemControllerbyCategoryId)
	eItem.GET("?keyword=item_name", controllers.GetItemControllerbyName)
	eItem.POST("", controllers.CreateItemController)
	eItem.PUT("/:id", controllers.UpdateItemController)
	eItem.DELETE("/:id", controllers.DeleteItemController)
	eItem.POST("/category", controllers.CreateCategoryController)
	eItem.GET("/category", controllers.GetCategoryControllerAll)
	eItem.GET("/category/:id", controllers.GetCategoryController)

	return e
}
