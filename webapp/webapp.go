package main

import(
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>indexHandler</h1>"))
}

func keyHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("<h1>keyHandler</h1>"))
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>loginHandler</h1>"))
}

func apiHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>apiHandler</h1>"))
}

func statisticsHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>statisticsHandler</h1>"))
}

func StartSrv(port string){
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/statistics", statisticsHandler)
	mux.HandleFunc("/key", keyHandler)
	http.ListenAndServe(":"+port, mux)
}

func main(){
	StartSrv("8080")
}