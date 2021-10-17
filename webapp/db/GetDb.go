package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "go-webapp/webapp/config"
)

var connection *sql.DB

func init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	                           config.GetDbHost(), config.GetDbPort(), config.GetDbUser(), config.GetDbPassword(), config.GetDbName())
    db, err := sql.Open("postgres", psqlconn)
    connection = db
    checkError(err)	
}

func GetDb() *sql.DB {
    return connection
}

func checkError(err error) {
	if err != nil {
	    panic(err)
	}
}