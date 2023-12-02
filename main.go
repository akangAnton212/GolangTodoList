package main

import (
	"github.com/akangAnton212/todolist/database"
	"github.com/akangAnton212/todolist/routes"
	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()

	if err != nil {
		panic(err)
	}
	e := echo.New()

	routes.RoutesApis(e)

	// e.GET("/todos", func(ctx echo.Context) error {
	// 	var jsonString = `{"Name": "john wick", "Age": 27}`
	// 	// var jsonData = []byte(jsonString)
	// 	responseApi := helpers.ResponseData(true, "Data Found", jsonString)

	// 	// Return the person as JSON
	// 	return ctx.JSON(http.StatusOK, responseApi)
	// })

	// e.POST("/todos", func(ctx echo.Context) error {
	// title := ctx.FormValue("title")
	// description := ctx.FormValue("description")

	// _, err := db.Exec("INSERT INTO todos (title, description, done) VALUES ( ?, ? , ?)", title, description, false)

	// if err != nil {
	// 	responseApi := helpers.ResponseData(false, "Internal Server Error", err.Error())
	// 	return ctx.JSON(http.StatusInternalServerError, responseApi)
	// }

	// // You can also retrieve other types of data like JSON
	// // var data map[string]interface{}
	// // if err := c.Bind(&data); err != nil {
	// // 	return err
	// // }

	// // Do something with the data
	// responseApi := helpers.ResponseData(true, "Create Data Success", "")
	// return ctx.JSON(http.StatusOK, responseApi)
	// })

	e.Start(":8080")
}
