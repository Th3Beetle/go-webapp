package main

import(
	"net/http"

	"go-webapp/webapp/handlers"
)

func statisticsHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>statisticsHandler</h1>"))
}

func StartSrv(port string){
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/api", handlers.ApiHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/statistics", handlers.StatisticsHandler)
	mux.HandleFunc("/key", handlers.KeyHandler)
	http.ListenAndServe(":"+port, mux)
}

func main(){
	StartSrv("8080")
}