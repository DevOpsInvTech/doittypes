package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/domains", homeHandler)
	r.HandleFunc("/hosts", homeHandler)
	r.HandleFunc("/groups", homeHandler)
	r.HandleFunc("/group_vars", homeHandler)
	r.HandleFunc("/group_vars/{group}/list", homeHandler)
	r.HandleFunc("/api/1/{type}/{name}/{action}", homeHandler).Methods("POST", "DELETE", "PUT", "GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
