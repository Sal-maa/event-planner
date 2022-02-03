package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func FetchConnection() *sql.DB {
	connection := "root:toor@tcp(172.17.0.1:3306)/event-planner?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	return db
}
