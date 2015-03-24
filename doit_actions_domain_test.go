package main

import "testing"

func TestDoitActionAddDomain(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddDomain.db")
	ds.Store.InitSchema(true)
	d, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetDomain(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Domain successfully created")
	ds.CloseDatastore()
}
