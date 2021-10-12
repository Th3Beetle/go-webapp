package handlers

import(
	"net/http"
)

func StatisticsHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>statisticsHandler</h1>"))
}