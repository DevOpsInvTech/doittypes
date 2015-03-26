package main

import (
	"encoding/json"
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
			h, err := ds.GetHostByName(d, reqName)
			if err != nil {
				panic(err)
			}
			data, err := json.Marshal(h)
			if err != nil {
				w.WriteHeader(500)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "POST":
			h, err := ds.AddHost(d, reqName)
			data, err := json.Marshal(h)
			if err != nil {
				w.WriteHeader(500)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "PUT":
		//TODO: Add host items here
		case "DELETE":
			h, err := ds.GetHostByName(d, reqName)
			if err != nil {
				panic(err)
			}
			err = ds.RemoveHost(d, h)
			if err != nil {
				w.WriteHeader(500)
			}
			w.WriteHeader(200)
		}
	case "group":
		if d.ID == 0 && d.Name == "" {
			panic("foo")
		}
		switch r.Method {
		case "GET":
			g, err := ds.GetGroupByName(d, reqName)
			if err != nil {
				panic(err)
			}
			data, err := json.Marshal(g)
			if err != nil {
				w.WriteHeader(500)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "POST":
			h, err := ds.AddGroup(d, reqName)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", h)
		case "PUT":
		//TODO: Add group items here
		case "DELETE":
			g, err := ds.GetGroupByName(d, reqName)
			if err != nil {
				panic(err)
			}
			err = ds.RemoveGroup(d, g)
			if err != nil {
				w.WriteHeader(500)
			}
			w.WriteHeader(200)
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
			data, err := json.Marshal(v)
			if err != nil {
				w.WriteHeader(500)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "POST":
			host, err := ds.AddVar(d, reqName, reqValue)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", host)
		case "PUT":
		case "DELETE":
			v, err := ds.GetVarByName(d, reqName)
			if err != nil {
				panic(err)
			}
			err = ds.RemoveVar(d, v)
			if err != nil {
				w.WriteHeader(500)
			}
			w.WriteHeader(200)
		}
	case "domain":
		switch r.Method {
		case "GET":
			d, err = ds.GetDomainByName(reqName)
			if err != nil {
				panic(err)
			}
			data, err := json.Marshal(d)
			if err != nil {
				w.WriteHeader(500)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		case "POST":
			domain, err := ds.AddDomain(reqName)
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", domain)
		case "PUT":
			//TODO: Add Domain items here
		case "DELETE":
			d, err := ds.GetDomainByName(reqName)
			if err != nil {
				panic(err)
			}
			err = ds.RemoveDomain(d)
			if err != nil {
				w.WriteHeader(500)
			}
			w.WriteHeader(200)
		}

	}
}
