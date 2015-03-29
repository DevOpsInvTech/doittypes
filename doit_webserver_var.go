package main

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiVarHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqValue := r.Form.Get("value")
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
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		data, err := json.Marshal(v)
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
		_, err := ds.AddVar(d, reqName, reqValue)
		if err != nil {
			//TODO: What error to throw here?
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	case "PUT":
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		err = ds.UpdateVar(d, v.ID, reqValue)
		if err != nil {
			//TODO: WHAT TO RETURN HERE?
			w.WriteHeader(http.StatusNotImplemented)
			ds.logger(r, http.StatusNotImplemented, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	case "DELETE":
		v, err := ds.GetVarByName(d, reqName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		err = ds.RemoveVar(d, v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ds.logger(r, http.StatusInternalServerError, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	}
}
