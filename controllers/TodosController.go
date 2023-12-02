package controllers

import (
	"github.com/akangAnton212/todolist/services"
	"github.com/labstack/echo"
)

func GetAllTodoList(ctx echo.Context) error {
	return services.GetAllTodoList(ctx)
}

func PostTodoList(ctx echo.Context) error {
	return services.PostTodoList(ctx)
}

func UpdateTodoList(ctx echo.Context) error {
	return services.UpdateTodoList(ctx)
}

func GetTodoById(ctx echo.Context) error {
	return services.GetTodoById(ctx)
}
