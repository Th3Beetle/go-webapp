package handlers

import(
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>loginHandler</h1>"))
}