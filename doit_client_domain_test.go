package main

import (
	"log"
	"testing"
)

func TestClientCreateDomain(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.CreateDomain(&Domain{Name: "domain1"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientGetDomain(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	d, err := dc.GetDomain(&Domain{Name: "domain1"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", d)
}

func TestClientGetDomains(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	d, err := dc.GetDomains()
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", d)
}

func TestClientDeleteDomain(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.DeleteDomain(&Domain{Name: "domain1"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = dc.GetDomain(&Domain{Name: "domain1"})
	if err != nil {
		t.Log("Domain removed succesfully")
	} else {
		t.Fatal("Domain found")
	}
}
