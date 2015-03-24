package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (ds *DoitServer) homeHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) domainHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) hostHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) groupHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) varsHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) ansibleHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) apiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domain := vars["type"]
	reqType := vars["type"]
	reqName := vars["name"]

	log.Println("Type: ", reqType)
	log.Println("Name: ", reqName)
	log.Println("Method: ", r.Method)
	log.Println("Domain: ", domain)

	switch reqType {
	case "host":
		switch r.Method {
		case "GET":
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
