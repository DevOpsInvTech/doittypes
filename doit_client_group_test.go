package main

import (
	"log"
	"testing"
)

func TestClientCreateGroup(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.CreateGroup(&Domain{Name: "foo"}, &Group{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientGetGroup(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	v, err := dc.GetGroup(&Domain{Name: "foo"}, &Group{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", v)
}

func testClientUpdateGroup(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.UpdateGroup(&Domain{Name: "foo"}, &Group{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = dc.GetGroup(&Domain{Name: "foo"}, &Group{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientDeleteGroup(t *testing.T) {
	dc := &DoitClient{}
	dc.SetURL("http://localhost:8080/api/1")
	err := dc.DeleteGroup(&Domain{Name: "foo"}, &Group{Name: "hello2"})
	if err != nil {
		t.Fatal(err)
	}
	v, err := dc.GetGroup(&Domain{Name: "foo"}, &Group{Name: "hello2"})
	if err != nil {
		t.Log("Group removed succesfully")
	} else {
		t.Fatal(v, "Group found")
	}
}
