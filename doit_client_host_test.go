package main

import (
	"log"
	"testing"
)

func TestClientCreateHost(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.CreateHost(&Domain{Name: "foo"}, &Host{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientGetHost(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	v, err := dc.GetHost(&Domain{Name: "foo"}, &Host{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", v)
}

func testClientUpdateHost(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.UpdateHost(&Domain{Name: "foo"}, &Host{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = dc.GetHost(&Domain{Name: "foo"}, &Host{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientDeleteHost(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.DeleteHost(&Domain{Name: "foo"}, &Host{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	v, err := dc.GetHost(&Domain{Name: "foo"}, &Host{Name: "hello2"})
	if err != nil {
		t.Log("Host removed succesfully")
	} else {
		t.Fatal(v, "Host found")
	}
}
