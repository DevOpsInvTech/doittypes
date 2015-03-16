package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/domains", domainHandler)
	r.HandleFunc("/hosts", hostHandler)
	r.HandleFunc("/groups", groupHandler)
	r.HandleFunc("/group_vars", varsHandler)
	r.HandleFunc("/group_vars/{group}/list", varsHandler)
	r.HandleFunc("/api/1/{type}/{name}", apiHandler).Methods("POST", "DELETE", "PUT", "GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
