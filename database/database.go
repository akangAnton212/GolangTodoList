package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbPublic *sql.DB

func InitDb() *sql.DB {
	dsn := "root@tcp(localhost:3306)/todolist"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	dbPublic = db
	return db
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	return dbPublic
}
