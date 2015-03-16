package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func domainHandler(w http.ResponseWriter, r *http.Request) {

}

func hostHandler(w http.ResponseWriter, r *http.Request) {

}

func groupHandler(w http.ResponseWriter, r *http.Request) {

}

func varsHandler(w http.ResponseWriter, r *http.Request) {

}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqType := vars["type"]
	reqName := vars["name"]

	log.Println("Type: ", reqType)
	log.Println("Name: ", reqName)
	log.Println("Method: ", r.Method)

	switch reqType {
	case "host":
		switch r.Method {
		case "GET":
		//Check for list
		case "PUT":
		case "POST":
		case "DELETE":
		}
	case "group":
		switch r.Method {
		case "GET":
		case "PUT":
		case "POST":
		case "DELETE":
		}
	case "domain":
		switch r.Method {
		case "GET":
		case "PUT":
		case "POST":
		case "DELETE":
		}
	case "var":
		switch r.Method {
		case "GET":
		case "PUT":
		case "POST":
		case "DELETE":
		}
	}
}
