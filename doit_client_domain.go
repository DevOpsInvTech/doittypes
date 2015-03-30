package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (dc *DoitClient) GetDomain(d *Domain) (*Domain, error) {
	newDomain := &Domain{}
	res, err := http.Get(dc.createAPIURL("domain", d.Name, "", ""))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &newDomain)
	if res.StatusCode == 200 {
		return newDomain, nil
	}
	return nil, errors.New(res.Status)
}

func (dc *DoitClient) GetDomains() ([]*Domain, error) {
	newDomains := []*Domain{}
	res, err := http.Get(dc.createAPIURL("domains", "", "", ""))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &newDomains)
	if res.StatusCode == 200 {
		return newDomains, nil
	}
	return nil, errors.New(res.Status)
}

func (dc *DoitClient) CreateDomain(d *Domain) error {
	res, err := http.Post(dc.createAPIURL("domain", d.Name, "", ""), JSONMime, nil)
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

func (dc *DoitClient) DeleteDomain(d *Domain) error {
	req, err := http.NewRequest("DELETE", dc.createAPIURL("domain", d.Name, "", ""), nil)
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
