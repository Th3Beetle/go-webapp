package handlers

import(
	"go-webapp/webapp/db"


	"net/http"

)

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	db := db.GetDb()
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

        _, err := db.Exec(`INSERT INTO users(login, password) VALUES($1, $2)`, login, password)
        if err != nil {
        	panic(err.Error())
        }

	}
}