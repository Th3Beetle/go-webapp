package handlers

import(
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>apiHandler</h1>"))
}