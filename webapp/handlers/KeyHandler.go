package handlers

import(
	"net/http"
)

func KeyHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("<h1>keyHandler</h1>"))
}