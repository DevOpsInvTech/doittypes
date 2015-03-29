package main

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiGroupHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqName := vars["name"]

	d := &Domain{}

	if len(domain) > 0 {
		var err error
		d, err = ds.GetDomainByName(domain)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
	}
	switch r.Method {
	case "GET":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		data, err := json.Marshal(g)
		if err != nil {
			log.Errorln("Unable to marshal json", data)
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		ds.logger(r, http.StatusOK, len(data))
	case "POST":
		_, err := ds.AddGroup(d, reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	case "PUT":
		//TODO: Add group items here
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
	case "DELETE":
		g, err := ds.GetGroupByName(d, reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		err = ds.RemoveGroup(d, g)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	}
}
