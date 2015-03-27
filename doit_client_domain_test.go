package main

import (
	"log"
	"testing"
)

func TestClientGetDomain(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	d, err := dc.GetDomain(&Domain{Name: "foo"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", d)
}
