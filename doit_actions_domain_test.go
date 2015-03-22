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

func TestDoitActionAddDomainVar(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddDomainVar.db")
	ds.Store.InitSchema(true)
	d, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddDomainVars(d.ID, Var{Name: "Var1", Value: "Val1"}); err != nil {
		t.Fatal(err)
	}
	cd, err := ds.GetDomain(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cd)
	if len(cd.Vars) == 1 {
		t.Log("One Var seen")
	} else {
		t.Fatal("No Var seen")
	}
	ds.CloseDatastore()
}
func TestDoitActionAddDomainVars(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddDomainVars.db")
	ds.Store.InitSchema(true)
	d, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddDomainVars(d.ID, []Var{Var{Name: "Var1", Value: "Val1"}, Var{Name: "Var2", Value: "Val2"}}...); err != nil {
		t.Fatal(err)
	}
	cd, err := ds.GetDomain(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cd)
	if len(cd.Vars) == 2 {
		t.Log("Two Hosts seen")
	} else {
		t.Fatal("No Host seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddDomainHost(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddDomainHost.db")
	ds.Store.InitSchema(true)
	d, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddDomainHosts(d.ID, Host{Name: "Host11"}); err != nil {
		t.Fatal(err)
	}
	cd, err := ds.GetDomain(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cd)
	if len(cd.Hosts) == 1 {
		t.Log("One Host seen")
	} else {
		t.Fatal("No Host seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddDomainHosts(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddDomainHosts.db")
	ds.Store.InitSchema(true)
	d, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddDomainHosts(d.ID, []Host{Host{Name: "Host11"}, Host{Name: "Host2"}}...); err != nil {
		t.Fatal(err)
	}
	cd, err := ds.GetDomain(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cd)
	if len(cd.Hosts) == 2 {
		t.Log("Two Hosts seen")
	} else {
		t.Fatal("No Hosts seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddDomainGroup(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddDomainGroup.db")
	ds.Store.InitSchema(true)
	d, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddDomainGroups(d.ID, Group{Name: "Group1"}); err != nil {
		t.Fatal(err)
	}
	cd, err := ds.GetDomain(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cd)
	if len(cd.Groups) == 1 {
		t.Log("One Group seen")
	} else {
		t.Fatal("No Group seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddDomainGroups(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddDomainGroups.db")
	ds.Store.InitSchema(true)
	d, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddDomainGroups(d.ID, []Group{Group{Name: "Group1"}, Group{Name: "Group2"}}...); err != nil {
		t.Fatal(err)
	}
	cd, err := ds.GetDomain(d.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cd)
	if len(cd.Groups) == 2 {
		t.Log("Two Groups seen")
	} else {
		t.Fatal("No Groups seen")
	}
	ds.CloseDatastore()
}
