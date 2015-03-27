package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (dc *DoitClient) GetVar(d *Domain, v *Var) (*Var, error) {
	newVar := &Var{}
	res, err := http.Get(dc.createAPIURL("var", v.Name, "", d.Name))
	if err != nil {
		return newVar, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return newVar, err
	}
	json.Unmarshal(data, &newVar)
	return newVar, nil
}

func (dc *DoitClient) CreateVar(d *Domain, v *Var) error {
	res, err := http.Post(dc.createAPIURL("var", v.Name, v.Value, d.Name), JSONMime, nil)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

func (dc *DoitClient) DeleteVar(d *Domain, v *Var) error {
	req, err := http.NewRequest("DELETE", dc.createAPIURL("var", v.Name, "", d.Name), nil)
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
	return nil
}
