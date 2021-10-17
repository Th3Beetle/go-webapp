package handlers

import(
	"net/http"
	"go-webapp/webapp/db"
	"database/sql"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	db := db.GetDb()
    if db == nil {
    	panic("panic")
    }
    getUser := `select login, password from users;`
    values := db.QueryRow(getUser)
    var login string
    var password string
    switch err := values.Scan(&login, &password); err {
    case sql.ErrNoRows:
    	fmt.Println("No rows")
    case nil:
    default:
    	panic(err)
    }    
	w.Write([]byte("<h1>indexHandler</h1>" + login + " " + password))
}