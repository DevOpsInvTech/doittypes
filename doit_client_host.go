package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (dc *DoitClient) GetHost(d *Domain, h *Host) (*Host, error) {
	newHost := &Host{}
	res, err := http.Get(dc.createAPIURL("host", h.Name, "", d.Name))
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return newHost, err
	}
	json.Unmarshal(data, &newHost)
	return newHost, nil
}
