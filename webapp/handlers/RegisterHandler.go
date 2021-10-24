package handlers

import(
	"go-webapp/webapp/db"
    "database/sql"

	"net/http"

)

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./html/register.html")

	case "POST":
        login := r.FormValue("login")
        password := r.FormValue("password")


        if login == "" || password == "" {
        	w.Write([]byte("<h1>login or password missing</h1>"))
            return
        }

        db := db.GetDb()

        if loginIsOccupied(login, db) {
        	w.Write([]byte("<h1>login is occupied</h1>"))
            return
        }

        _, err := db.Exec(`INSERT INTO users(login, password) VALUES($1, $2)`, login, password)
        if err != nil {
        	panic(err.Error())
        }
        w.Write([]byte("<h1>Account created</h1>"))

	}
}

func loginIsOccupied(login string, db *sql.DB) bool {
	getUser := `select login from users where login = $1;`
    row := db.QueryRow(getUser, login)
    if row.Scan() == sql.ErrNoRows {
    	return false
    }
    return true
}