package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDB *sql.DB
	err     error
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	// SERVER   = "localhost"
	SERVER   = "host.docker.internal"
	PORT     = "3306"
	DATABASE = "authorization"
)

func ConnectDB() {
	conectionString := USERNAME + ":" + PASSWORD + "@" + NETWORK + "(" + SERVER + ":" + PORT + ")/" + DATABASE + "?parseTime=true"
	MysqlDB, err = sql.Open("mysql", conectionString)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database:", DATABASE)
	}

}
