package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (dc *DoitClient) GetGroup(d *Domain, g *Group) (*Group, error) {
	newGroup := &Group{}
	res, err := http.Get(dc.createAPIURL("group", g.Name, "", d.Name))
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return newGroup, err
	}
	json.Unmarshal(data, &newGroup)
	return newGroup, nil
}
