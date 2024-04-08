package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	PostgreDB *sql.DB
)

func ConnectPostgreDB() {

	var USERNAME = "root"
	var PASSWORD = "root"
	var SERVER = "host.docker.internal"
	var PORT = "5432"
	var DATABASE = "legacy"
	var err error

	conectionString := "user=" + USERNAME + " dbname=" + DATABASE + "password=" + PASSWORD + " host=" + SERVER + " port=" + PORT + " sslmode=disable"
	PostgreDB, err = sql.Open("postgres", conectionString)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database:", DATABASE)
	}

}
