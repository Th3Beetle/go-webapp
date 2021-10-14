package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "go_test"
    password = "test"
    dbname = "go_test"
)

func GetDb() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	                           host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlconn)
    checkError(err)

    return db	
}

func checkError(err error) {
	if err != nil {
	    panic(err)
	}
}