package main

import (
	"log"
	"testing"
)

func TestClientGetHost(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	h, err := dc.GetHost(&Domain{Name: "foo"}, &Host{Name: "food"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", h)
}
