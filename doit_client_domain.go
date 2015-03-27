package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (dc *DoitClient) GetDomain(d *Domain) (*Domain, error) {
	newDomain := &Domain{}
	res, err := http.Get(dc.createAPIURL("domain", d.Name, "", ""))
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return newDomain, err
	}
	json.Unmarshal(data, &newDomain)
	return newDomain, nil
}
