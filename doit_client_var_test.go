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

func TestClientGetVars(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	v, err := dc.GetVars(&Domain{Name: "foo"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", v)
}

func TestClientUpdateVar(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.UpdateVar(&Domain{Name: "foo"}, &Var{Name: "hello2", Value: "over"})
	if err != nil {
		t.Fatal(err)
	}
	v, err := dc.GetVar(&Domain{Name: "foo"}, &Var{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	if v.Value != "over" {
		t.Fatal("Var not updated")
	}
}

func TestClientDeleteVar(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.DeleteVar(&Domain{Name: "foo"}, &Var{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = dc.GetVar(&Domain{Name: "foo"}, &Var{Name: "hello2"})
	if err != nil {
		t.Log("Var removed succesfully")
	} else {
		t.Fatal("Var found")
	}
}
