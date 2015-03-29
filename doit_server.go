package main

import (
	"net"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

//DoitServer A webserver to frontend a DOIT database
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

//Listen Starts the webserver for listening
func (ds *DoitServer) Listen(port *string, config *DoitConfig) (err error) {
	err = ds.OpenDatastore(config.Storage.Type, config.Storage.Location)
	if err != nil {
		return err
	}
	ds.Store.InitSchema(false)
	r := mux.NewRouter()
	r.HandleFunc("/", ds.homeHandler)
	r.HandleFunc("/api/1/domain/{name}", ds.apiDomainHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/var/{name}", ds.apiVarHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/group/{name}", ds.apiGroupHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/group/{name}/var/{varName}", ds.apiGroupHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/group/{name}/vars", ds.apiGroupHandler).Methods("GET")
	r.HandleFunc("/api/1/group/{name}/host/{hostName}", ds.apiGroupHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/group/{name}/hosts", ds.apiGroupHandler).Methods("GET")
	r.HandleFunc("/api/1/host/{name}", ds.apiHostHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/host/{name}/var/{varName}", ds.apiHostHandler).Methods("POST", "DELETE", "PUT", "GET")
	r.HandleFunc("/api/1/host/{name}/vars", ds.apiHostHandler).Methods("GET")
	r.HandleFunc("/api/ansible/domain/{name}", ds.ansibleHandler).Methods("GET")

	http.Handle("/", r)

	log.Infoln("Staring webserver")
	if err := http.ListenAndServe(net.JoinHostPort("", *port), nil); err != nil {
		log.Errorln(err)
		return err
	}
	return err
}
