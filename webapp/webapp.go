package main

import(
	"net/http"

	"go-webapp/webapp/handlers"
)


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