package handlers

import(
	"net/http"
    "database/sql"

    "go-webapp/webapp/db"
	"go-webapp/webapp/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":{
		http.ServeFile(w, r, "./html/login.html")
	}
	case "POST": {
		login := r.FormValue("login")
        password := r.FormValue("password")

        if login == "" || password == "" {
        	w.Write([]byte("<h1>login or password missing</h1>"))
            return
        }

        db := db.GetDb()

        if !userExists(login, password, db) {
        	w.Write([]byte("<h1>Login or password invalid</h1>"))
            return
        }

        var token string
        cookie, err := r.Cookie("gsession")
        if err != nil {
        	panic("Cookie missing")
        } else {
        	token = cookie.Value
        }
        session, err := sessions.GetSession(token)
        if err != nil {
        	panic("missing session")
        }
        session.Is_Authenticated = true        

        w.Write([]byte("<h1>Authenticated</h1>" + token))
        return
    }
	}
}


func userExists(login string, password string, db *sql.DB) bool {
	getUser := `select login from users where login = $1 and password = $2;`
    row := db.QueryRow(getUser, login, password)
    if row.Scan() == sql.ErrNoRows {
    	return false
    }
    return true
}