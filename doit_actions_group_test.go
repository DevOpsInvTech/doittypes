package main

import "testing"

func TestDoitActionAddGroup(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroup.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	g, err := ds.AddGroup(domain, "Group1")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetGroup(domain, g.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Group successfully created")
	ds.CloseDatastore()
}

func TestDoitActionAddGroupVar(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroupVar.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	g, err := ds.AddGroup(domain, "Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupVars(domain, g.ID, Var{Name: "Var1", Value: "Val1", Domain: domain}); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(domain, g.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cg)
	if len(cg.Vars) == 1 {
		t.Log("One Var seen")
	} else {
		t.Fatal("No Var seen")
	}
	ds.CloseDatastore()
}
func TestDoitActionAddGroupVars(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroupVars.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	g, err := ds.AddGroup(domain, "Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupVars(domain, g.ID, []Var{Var{Name: "Var1", Value: "Val1", Domain: domain}, Var{Name: "Var2", Value: "Val2", Domain: domain}}...); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(domain, g.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cg)
	if len(cg.Vars) == 2 {
		t.Log("Two Hosts seen")
	} else {
		t.Fatal("No Host seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddGroupHost(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroupHost.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	g, err := ds.AddGroup(domain, "Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupHosts(domain, g.ID, Host{Name: "Host1", Domain: domain}); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(domain, g.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cg)
	if len(cg.Hosts) == 1 {
		t.Log("One Host seen")
	} else {
		t.Fatal("No Host seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddGroupHosts(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroupHosts.db")
	ds.Store.InitSchema(true)
	domain, err := ds.AddDomain("Domain1")
	if err != nil {
		t.Fatal(err)
	}
	g, err := ds.AddGroup(domain, "Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupHosts(domain, g.ID, []Host{Host{Name: "Host11", Domain: domain}, Host{Name: "Host2", Domain: domain}}...); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(domain, g.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cg)
	if len(cg.Hosts) == 2 {
		t.Log("Two Hosts seen")
	} else {
		t.Fatal("No Hosts seen")
	}
	ds.CloseDatastore()
}
