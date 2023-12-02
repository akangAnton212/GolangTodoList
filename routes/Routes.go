package routes

import (
	"github.com/akangAnton212/todolist/controllers"
	"github.com/labstack/echo"
)

// RegisterRoutes is a function that registers all routes for the Echo application.
func RoutesApis(e *echo.Echo) {
	e.GET("/todos", controllers.GetAllTodoList)
	e.POST("/todos", controllers.PostTodoList)
	e.POST("/todos/update", controllers.UpdateTodoList)
	e.GET("/todos/:id", controllers.GetTodoById)
}
