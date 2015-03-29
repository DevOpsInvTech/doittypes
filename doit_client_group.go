package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (dc *DoitClient) GetGroup(d *Domain, g *Group) (*Group, error) {
	newGroup := &Group{}
	res, err := http.Get(dc.createAPIURL("group", g.Name, "", d.Name))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &newGroup)
	if res.StatusCode == 200 {
		return newGroup, nil
	}
	return nil, errors.New(res.Status)
}

func (dc *DoitClient) CreateGroup(d *Domain, g *Group) error {
	res, err := http.Post(dc.createAPIURL("group", g.Name, "", d.Name), JSONMime, nil)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return errors.New(res.Status)
}

//UpdateGroup d
//NOT IMPLEMENTED
func (dc *DoitClient) UpdateGroup(d *Domain, g *Group) error {
	req, err := http.NewRequest("PUT", dc.createAPIURL("group", g.Name, "", d.Name), nil)
	req.Header.Add("Content-Type", JSONMime)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return errors.New(res.Status)
}

func (dc *DoitClient) DeleteGroup(d *Domain, g *Group) error {
	req, err := http.NewRequest("DELETE", dc.createAPIURL("group", g.Name, "", d.Name), nil)
	req.Header.Add("Content-Type", JSONMime)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return errors.New(res.Status)
}
