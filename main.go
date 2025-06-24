package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josetorresdeveloper/API-REST-GO25/routes"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func main() {
	//fmt.Println("Hello world!")
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
