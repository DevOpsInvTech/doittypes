package main

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiDomainHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	vars := mux.Vars(r)
	reqName := vars["name"]

	switch r.Method {
	case "GET":
		d, err := ds.GetDomainByName(reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		data, err := json.Marshal(d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		ds.logger(r, http.StatusOK, len(data))
	case "POST":
		_, err := ds.AddDomain(reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	case "PUT":
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
	case "DELETE":
		d, err := ds.GetDomainByName(reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		err = ds.RemoveDomain(d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	}
}

func (ds *DoitServer) apiDomainsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}

	switch r.Method {
	case "GET":
		d, err := ds.GetDomains()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		data, err := json.Marshal(d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		ds.logger(r, http.StatusOK, len(data))
	}
}
