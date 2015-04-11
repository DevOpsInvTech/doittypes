package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func (ds *DoitServer) DomainCheck(dName string) (d *Domain, err error) {
	if len(dName) > 0 {
		var err error
		d, err = ds.GetDomainByName(dName)
		if err != nil {
			return nil, err
		}
		return d, nil
	}
	return nil, errors.New("Domain string not valid")
}

func (ds *DoitServer) homeHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) ansibleHandler(w http.ResponseWriter, r *http.Request) {

}

func (ds *DoitServer) logger(r *http.Request, status int, retSize int) {
	t := time.Now()
	zone, _ := t.Zone()
	log.Infof("%s %s %s [%s] \"%s %s %s\" %d %d", r.RemoteAddr, "-", "-", fmt.Sprintf("%d/%s/%d:%d:%d:%d %s", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(), zone), r.Method, r.URL.RequestURI(), r.Proto, status, retSize)
}

func (ds *DoitServer) oldApiHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorln("Unable to parse message", err)
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	domain := r.Form.Get("domain")
	reqValue := r.Form.Get("value")
	reqType := vars["type"]
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

	switch reqType {
	case "host":
		switch r.Method {
		case "GET":
			h, err := ds.GetHostByName(d, reqName)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				ds.logger(r, http.StatusNotFound, 0)
				return
			}
			data, err := json.Marshal(h)
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
		}
	case "group":
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
	case "var":
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
	case "domain":
		switch r.Method {
		case "GET":
			d, err = ds.GetDomainByName(reqName)
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
}
