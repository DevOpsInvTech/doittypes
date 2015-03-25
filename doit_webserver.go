package main

import (
	"fmt"
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
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqValue := r.Form.Get("value")
	reqType := vars["type"]
	reqName := vars["name"]

	log.Println("Type: ", reqType)
	log.Println("Name: ", reqName)
	log.Println("Method: ", r.Method)
	log.Println("Domain: \"", domain, "\"", len(domain))

	d := &Domain{}

	if len(domain) > 0 {
		var err error
		d, err = ds.GetDomainByName(domain)
		if err != nil {
			panic(err)
		}
	}

	switch reqType {
	case "host":
		if d.ID == 0 && d.Name == "" {
			panic("foo")
		}
		switch r.Method {
		case "GET":
			host, err := ds.GetHostByName(d, reqName)
			if err != nil {
				panic(err)
			}
			fmt.Println(host)
		case "PUT":
			host, err := ds.AddHost(d, reqName)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", host)
		case "POST":
		case "DELETE":
		}
	case "group":
		if d.ID == 0 && d.Name == "" {
			panic("foo")
		}
		switch r.Method {
		case "GET":
		case "PUT":
			host, err := ds.AddGroup(d, reqName)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", host)
		case "POST":
		case "DELETE":
		}
	case "var":
		if d.ID == 0 && d.Name == "" {
			panic("foo")
		}
		switch r.Method {
		case "GET":
			v, err := ds.GetVarByName(d, reqName)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", v)
		case "PUT":
			host, err := ds.AddVar(d, reqName, reqValue)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", host)
		case "POST":
		case "DELETE":
		}
	case "domain":
		switch r.Method {
		case "GET":
			d, err = ds.GetDomainByName(domain)
			if err != nil {
				panic(err)
			}
		case "PUT":
			domain, err := ds.AddDomain(reqName)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", domain)
		case "POST":
		case "DELETE":
		}

	}
}
