package main

import "testing"

func TestDoitActionAddHost(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddHost.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	newHost, err := ds.AddHost(domain, "Steve")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetHost(domain, newHost.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Host successfully created")
	ds.CloseDatastore()
}

func TestDoitActionAddHostVar(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddHostVar.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	newHost, err := ds.AddHost(domain, "Steve")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddHostVars(domain, newHost.ID, HostVar{Name: "Var1", Value: "Val1", Domain: domain}); err != nil {
		t.Fatal(err)
	}
	checkHost, err := ds.GetHost(domain, newHost.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(checkHost)
	if len(checkHost.Vars) == 1 {
		t.Log("One Var seen")
	} else {
		t.Fatal("No Var seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddHostVars(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddHostVars.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	newHost, err := ds.AddHost(domain, "Steve")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddHostVars(domain, newHost.ID, []HostVar{HostVar{Name: "Var1", Value: "Val1"}, HostVar{Name: "Var2", Value: "Val2"}}...); err != nil {
		t.Fatal(err)
	}
	checkHost, err := ds.GetHost(domain, newHost.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(checkHost)
	if len(checkHost.Vars) == 2 {
		t.Log("Two Vars seen")
	} else {
		t.Fatal("No Vars seen")
	}
	ds.CloseDatastore()
}

//Remove Host Tests

func TestDoitActionRemoveHost(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionRemoveHost.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	h, err := ds.AddHost(domain, "Steve")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.RemoveHost(domain, h.ID); err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetHost(domain, h.ID)
	if err == nil {
		t.Fatal("Host found in database")
	}
	t.Log("Host not found in database")
	ds.CloseDatastore()
}

func TestDoitActionRemoveHostAndVars(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionRemoveHostAndVars.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	h, err := ds.AddHost(domain, "Steve")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddHostVars(domain, h.ID, HostVar{Name: "Var1", Value: "Val1"}); err != nil {
		t.Fatal(err)
	}
	if err := ds.RemoveHost(domain, h.ID); err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetHost(domain, h.ID)
	if err == nil {
		t.Fatal("Host found in database")
	}
	t.Log("Host not found in database")
	ds.CloseDatastore()
}

func TestDoitActionRemoveHostVars(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionRemoveHostVars.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	h, err := ds.AddHost(domain, "Steve")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddHostVars(domain, h.ID, HostVar{Name: "Var1", Value: "Val1"}); err != nil {
		t.Fatal(err)
	}
	if err := ds.RemoveHostVars(domain, h.ID, HostVar{ID: 1, Name: "Var1", Value: "Val1"}); err != nil {
		t.Fatal(err)
	}
	checkHost := &Host{}
	checkHost, err = ds.GetHost(domain, h.ID)
	if err != nil {
		t.Fatal("Host not found in database")
	}
	if len(checkHost.Vars) > 0 {
		t.Fatal("Host vars found in host")
	}

	t.Log("Host vars not found in database")
	ds.CloseDatastore()
}
