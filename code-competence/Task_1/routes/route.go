package routes

import (
	"Task_1/controllers"
	m "Task_1/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

//	eUser := e.Group("/user")
//	eUser.POST("/register", controllers.RegisterUserController)
//	eUser.POST("/login", controllers.LoginUserController)
//	eUser.GET("/:id", controllers.GetUserProfile)
//	eUser.PUT("/:id", controllers.UpdateUserController)
//	eUser.DELETE("/:id", controllers.DeleteUserController)
	eItem := e.Group("/item")
	eItem.GET("", controllers.GetItemControllerAll)
	eItem.GET("/:id", controllers.GetItemById)
	eItem.GET("/category/:id", controllers.GetItemByCategori)
	eItem.GET("?keyword=item_name", controllers.GetItemByCategori)
	eItem.POST("", controllers.CreateItemController)
	eItem.PUT("/:id", controllers.UpdateItemController)
	eItem.DELETE("/:id", controllers.DeleteItemController)
	eItem.GET("/category", controllers.GetCategoryControllerAll)
	eItem.GET("/category/:id", controllers.GetCategoryController)
	

	return e
}
