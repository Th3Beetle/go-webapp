package main

import(
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func StartSrv(port string){
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}

func main(){
	StartSrv("8080")
}