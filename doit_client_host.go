package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (dc *DoitClient) GetHost(d *Domain, h *Host) (*Host, error) {
	newHost := &Host{}
	res, err := http.Get(dc.createAPIURL("host", h.Name, "", d.Name))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &newHost)
	if res.StatusCode == 200 {
		return newHost, nil
	}
	return nil, errors.New(res.Status)
}

func (dc *DoitClient) CreateHost(d *Domain, h *Host) error {
	res, err := http.Post(dc.createAPIURL("host", h.Name, "", d.Name), JSONMime, nil)
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

//UpdateHost s
//NOT IMPLEMENTED
func (dc *DoitClient) UpdateHost(d *Domain, h *Host) error {
	req, err := http.NewRequest("PUT", dc.createAPIURL("host", h.Name, "", d.Name), nil)
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

func (dc *DoitClient) DeleteHost(d *Domain, h *Host) error {
	req, err := http.NewRequest("DELETE", dc.createAPIURL("host", h.Name, "", d.Name), nil)
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
