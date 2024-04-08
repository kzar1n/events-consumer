package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDB *sql.DB
)

func ConnectMysqlDB() {

	var USERNAME = "root"
	var PASSWORD = "root"
	var SERVER = "host.docker.internal"
	var PORT = "3306"
	var DATABASE = "authorization"
	var err error

	conectionString := USERNAME + ":" + PASSWORD + "@tcp(" + SERVER + ":" + PORT + ")/" + DATABASE + "?parseTime=true"
	MysqlDB, err = sql.Open("mysql", conectionString)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database:", DATABASE)
	}

}
