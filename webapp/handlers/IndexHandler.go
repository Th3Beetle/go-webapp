package handlers

import(
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>indexHandler</h1>"))
}