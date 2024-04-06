package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func InitializeDatabase() {
	DB, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/authorization")
	if err != nil {
		panic(err)
	}
	defer DB.Close()
}
