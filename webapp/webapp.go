package main

import(
	"net/http"
	"strconv"

	"go-webapp/webapp/handlers"
	"go-webapp/webapp/config"
	"go-webapp/webapp/wrappers"
)


func StartSrv(port int){
	mux := http.NewServeMux()
    
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/api", handlers.ApiHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/statistics", handlers.StatisticsHandler)
	mux.HandleFunc("/key", handlers.KeyHandler)

	wrappedMux := wrappers.PreprocessRequest(mux)
	http.ListenAndServe(":"+strconv.Itoa(port), wrappedMux)
}

func main(){
	port := config.GetPort()
	StartSrv(port)
}