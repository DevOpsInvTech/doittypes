package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) apiHostVarHandler(w http.ResponseWriter, r *http.Request) {
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
	varName := vars["varName"]
	value := vars["value"]

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		hv, err := ds.GetHostVarByName(d, h, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(hv, w, r)
		if err != nil {
			return
		}
	case "POST":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		err = ds.AddHostVars(d, h.ID, &HostVar{Name: varName, Value: value, Domain: d})
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	case "PUT":
		//TODO: Add host items here
		log.Println(value)
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
	case "DELETE":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			w.WriteHeader(404)
		}
		err = ds.RemoveHostVars(d, h.ID, &HostVar{Name: varName, Value: value, Domain: d})
		if err != nil {
			//TODO: What error to throw here?
			ds.ReturnNotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
		return
	}
}

func (ds *DoitServer) apiHostHandler(w http.ResponseWriter, r *http.Request) {
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

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(h, w, r)
		if err != nil {
			return
		}
	case "POST":
		_, err := ds.AddHost(d, reqName)
		if err != nil {
			//TODO: What error to throw here?
			w.WriteHeader(http.StatusNotFound)
			ds.logger(r, http.StatusNotFound, 0)
			return
		}
		w.WriteHeader(http.StatusOK)
		ds.logger(r, http.StatusOK, 0)
	case "PUT":
		//TODO: Add host items here
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
	case "DELETE":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			w.WriteHeader(404)
		}
		err = ds.RemoveHost(d, h)
		if err != nil {
			w.WriteHeader(500)
		}
		w.WriteHeader(200)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
		return
	}
}

func (ds *DoitServer) apiHostsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(http.StatusInternalServerError)
		ds.logger(r, http.StatusInternalServerError, 0)
		return
	}
	domain := r.Form.Get("domain")

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostsByDomain(d)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(h, w, r)
		if err != nil {
			return
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
		return
	}
}

func (ds *DoitServer) apiHostVarsHandler(w http.ResponseWriter, r *http.Request) {
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

	d, err := ds.DomainCheck(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ds.logger(r, http.StatusBadRequest, 0)
		return
	}

	switch r.Method {
	case "GET":
		h, err := ds.GetHostByName(d, reqName)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		hv, err := ds.GetHostVars(d, h)
		if err != nil {
			ds.ReturnNotFound(w, r)
			return
		}
		ds.ReturnJSON(hv, w, r)
		if err != nil {
			return
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		ds.logger(r, http.StatusNotImplemented, 0)
		return
	}
}
