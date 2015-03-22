package main

import "testing"

func TestDoitActionAddGroup(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroup.db")
	ds.Store.InitSchema(true)
	g, err := ds.AddGroup("Group1")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ds.GetGroup(g.ID)
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
	g, err := ds.AddGroup("Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupVars(g.ID, Var{Name: "Var1", Value: "Val1"}); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(g.ID)
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
	g, err := ds.AddGroup("Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupVars(g.ID, []Var{Var{Name: "Var1", Value: "Val1"}, Var{Name: "Var2", Value: "Val2"}}...); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(g.ID)
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
	g, err := ds.AddGroup("Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupHosts(g.ID, Host{Name: "Host11"}); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(g.ID)
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
	g, err := ds.AddGroup("Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupHosts(g.ID, []Host{Host{Name: "Host11"}, Host{Name: "Host2"}}...); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(g.ID)
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

func TestDoitActionAddGroupDomain(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroupDomain.db")
	ds.Store.InitSchema(true)
	g, err := ds.AddGroup("Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupDomains(g.ID, Domain{Name: "Domain1"}); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(g.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cg)
	if len(cg.Domains) == 1 {
		t.Log("One Domain seen")
	} else {
		t.Fatal("No Domain seen")
	}
	ds.CloseDatastore()
}

func TestDoitActionAddGroupDomains(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitActionAddGroupDomains.db")
	ds.Store.InitSchema(true)
	g, err := ds.AddGroup("Group1")
	if err != nil {
		t.Fatal(err)
	}
	if err := ds.AddGroupDomains(g.ID, []Domain{Domain{Name: "Domain1"}, Domain{Name: "Domain2"}}...); err != nil {
		t.Fatal(err)
	}
	cg, err := ds.GetGroup(g.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cg)
	if len(cg.Domains) == 2 {
		t.Log("Two Domains seen")
	} else {
		t.Fatal("No Domains seen")
	}
	ds.CloseDatastore()
}
