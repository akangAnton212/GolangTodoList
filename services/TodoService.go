package services

import (
	"database/sql"
	"net/http"

	"github.com/akangAnton212/todolist/database"
	"github.com/akangAnton212/todolist/helpers"
	"github.com/labstack/echo"
)

func GetAllTodoList(ctx echo.Context) error {
	dbInstance := database.GetDB()
	rows, err := dbInstance.Query("SELECT * FROM todos")
	if err != nil {
		responseApi := helpers.ResponseData(false, "Internal Server Error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, responseApi)
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		responseApi := helpers.ResponseData(false, "Internal Server Error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, responseApi)
	}

	var result []map[string]interface{}
	for rows.Next() {
		// Create a slice of interface{} to store column values
		values := make([]interface{}, len(columns))
		// Create a slice of sql.RawBytes to represent each column's value
		rawResult := make([]sql.RawBytes, len(columns))

		// Assign pointers to each column's value to values slice
		for i := range values {
			values[i] = &rawResult[i]
		}

		// Scan the row into the values slice
		if err := rows.Scan(values...); err != nil {
			panic(err)
		}

		// Create a map to store the column name and value pairs
		rowData := make(map[string]interface{})
		for i, col := range columns {
			// Use string() conversion for sql.RawBytes to get the actual value
			rowData[col] = string(rawResult[i])
		}

		// Append the map to the result slice
		result = append(result, rowData)
	}
	if result != nil {
		responseApi := helpers.ResponseData(true, "Data Found", result)
		return ctx.JSON(http.StatusOK, responseApi)
	} else {
		responseApi := helpers.ResponseData(false, "Data Not Found", "")
		return ctx.JSON(http.StatusOK, responseApi)
	}
}

func PostTodoList(ctx echo.Context) error {
	dbInstance := database.GetDB()
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")

	_, err := dbInstance.Exec("INSERT INTO todos (title, description, done) VALUES ( ?, ? , ?)", title, description, false)

	if err != nil {
		responseApi := helpers.ResponseData(false, "Internal Server Error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, responseApi)
	}

	// You can also retrieve other types of data like JSON
	// var data map[string]interface{}
	// if err := c.Bind(&data); err != nil {
	// 	return err
	// }

	// Do something with the data
	responseApi := helpers.ResponseData(true, "Create Data Success", "")
	return ctx.JSON(http.StatusOK, responseApi)
}

func UpdateTodoList(ctx echo.Context) error {
	dbInstance := database.GetDB()
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	id := ctx.FormValue("id")
	done := ctx.FormValue("done")

	_, err := dbInstance.Exec("UPDATE todos SET title=?,description=?, done=? WHERE id=?", title, description, done, id)

	if err != nil {
		responseApi := helpers.ResponseData(false, "Internal Server Error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, responseApi)
	}

	// You can also retrieve other types of data like JSON
	// var data map[string]interface{}
	// if err := c.Bind(&data); err != nil {
	// 	return err
	// }

	// Do something with the data
	responseApi := helpers.ResponseData(true, "Update Data Success", "")
	return ctx.JSON(http.StatusOK, responseApi)
}

func GetTodoById(ctx echo.Context) error {
	dbInstance := database.GetDB()
	id := ctx.Param("id")
	rows, err := dbInstance.Query("SELECT * FROM todos where id=?", id)
	if err != nil {
		responseApi := helpers.ResponseData(false, "Internal Server Error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, responseApi)
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		responseApi := helpers.ResponseData(false, "Internal Server Error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, responseApi)
	}

	var result []map[string]interface{}
	for rows.Next() {
		// Create a slice of interface{} to store column values
		values := make([]interface{}, len(columns))
		// Create a slice of sql.RawBytes to represent each column's value
		rawResult := make([]sql.RawBytes, len(columns))

		// Assign pointers to each column's value to values slice
		for i := range values {
			values[i] = &rawResult[i]
		}

		// Scan the row into the values slice
		if err := rows.Scan(values...); err != nil {
			panic(err)
		}

		// Create a map to store the column name and value pairs
		rowData := make(map[string]interface{})
		for i, col := range columns {
			// Use string() conversion for sql.RawBytes to get the actual value
			rowData[col] = string(rawResult[i])
		}

		// Append the map to the result slice
		result = append(result, rowData)
	}

	if result != nil {
		responseApi := helpers.ResponseData(true, "Data Found", result)
		return ctx.JSON(http.StatusOK, responseApi)
	} else {
		responseApi := helpers.ResponseData(false, "Data Not Found", "")
		return ctx.JSON(http.StatusOK, responseApi)
	}
}
