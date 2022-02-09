package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func FetchConnection() *sql.DB {
	connection := "admin:password@tcp(database-1.c9tz5kkyxqmv.ap-northeast-3.rds.amazonaws.com:3306)/eplanner?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	return db
}
