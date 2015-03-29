package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type DoitServer struct {
	Store *DoitStorage
}

//OpenDatastore open datastore for writing
func (ds *DoitServer) OpenDatastore(t string, loc string) (err error) {
	s, err := NewStorage(t, loc)
	ds.Store = s
	return err
}

//CloseDatastore close datastore
func (ds *DoitServer) CloseDatastore() error {
	err := ds.Store.Close()
	return err
}

func (ds *DoitServer) Listen(port *string, config *DoitConfig) (err error) {
	err = ds.OpenDatastore(config.Storage.Type, config.Storage.Location)
	if err != nil {
		return err
	}
	ds.Store.InitSchema(true)
	r := mux.NewRouter()
	r.HandleFunc("/", ds.homeHandler)
	r.HandleFunc("/domains", ds.domainHandler)
	r.HandleFunc("/hosts", ds.hostHandler)
	r.HandleFunc("/groups", ds.groupHandler)
	r.HandleFunc("/group_vars", ds.varsHandler)
	r.HandleFunc("/group_vars/{group}/list", ds.varsHandler)
	r.HandleFunc("/api/1/{type}/{name}", ds.apiHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/domain/{name}", ds.apiHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/var/{name}", ds.apiHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/group/{name}", ds.apiHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/group/{name}/var/{varName}", ds.apiHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/group/{name}/host/{hostName}", ds.apiHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/host/{name}", ds.apiHostHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/host/{name}/var/{varName}", ds.apiHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/host/{name}/vars", ds.apiHandler).Methods("GET")
	r.HandleFunc("/api/ansible", ds.ansibleHandler).Methods("GET")

	http.Handle("/", r)

	if err := http.ListenAndServe(net.JoinHostPort("", *port), nil); err != nil {
		log.Println(err)
		return err
	}
	return err
}
