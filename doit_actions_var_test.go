package main

import "testing"

func TestDoitActionAddVar(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddVar.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	v, err := ds.AddVar(domain, "Val1", "Var1")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetVar(domain, v.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Var successfully created")
	ds.CloseDatastore()
}

func TestDoitActionRemoveVar(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionRemoveVar.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	v, err := ds.AddVar(domain, "Val1", "Var1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.RemoveVar(domain, v); err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetVar(domain, v.ID)
	if err == nil {
		t.Fatal("Var found in database")
	}
	t.Log("Var not found in database")
	ds.CloseDatastore()
}

func TestDoitActionUpdateVar(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionUpdateVar.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	v, err := ds.AddVar(domain, "Val1", "Var1")
	if err != nil {
		t.Fatal(err)
	}
	err = ds.UpdateVar(domain, v.ID, "Var2")
	if err != nil {
		t.Fatal(err)
	}
	cv, err := ds.GetVar(domain, v.ID)
	if err != nil {
		t.Fatal(err)
	}
	if cv.Value != "Var2" {
		t.Fatal("Var not updated")
	}
	t.Log("Var successfully updated")
	ds.CloseDatastore()
}
