package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("p", "8080", "Port")
	config := flag.String("c", "", "Load config file")
	serverMode := flag.Bool("s", false, "Enable server mode")

	flag.Parse()

	log.Println(port, serverMode, config)

	if *serverMode {
		r := mux.NewRouter()
		r.HandleFunc("/", homeHandler)
		r.HandleFunc("/domains", domainHandler)
		r.HandleFunc("/hosts", hostHandler)
		r.HandleFunc("/groups", groupHandler)
		r.HandleFunc("/group_vars", varsHandler)
		r.HandleFunc("/group_vars/{group}/list", varsHandler)
		r.HandleFunc("/api/1/{type}/{name}", apiHandler).Methods("POST", "DELETE", "PUT", "GET")

		http.Handle("/", r)
		if err := http.ListenAndServe(net.JoinHostPort("", *port), nil); err != nil {
			log.Println(err)
		}
	} else {
		//Act as client
	}
}
