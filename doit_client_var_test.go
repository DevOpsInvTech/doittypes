package main

import (
	"log"
	"testing"
)

func TestClientCreateVar(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.CreateVar(&Domain{Name: "foo"}, &Var{Name: "hello2", Value: "there"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientGetVar(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	v, err := dc.GetVar(&Domain{Name: "foo"}, &Var{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", v)
}

func TestClientDeleteVar(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.DeleteVar(&Domain{Name: "foo"}, &Var{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	v, err := dc.GetVar(&Domain{Name: "foo"}, &Var{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", v)
}
