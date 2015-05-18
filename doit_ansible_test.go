package doittypes

import "testing"

func TestDomainAnsibleBasic1(t *testing.T) {
	d := &Domain{Name: "Fool",
		Vars:   []*Var{&Var{Name: "VName", Value: "deepvalue"}},
		Groups: []*Group{&Group{Name: "GName"}},
		Hosts: []*Host{&Host{Name: "HName",
			Vars: []*Var{&Var{Name: "Hello", Value: "There"}}}}}
	f := d.MarshalAnsible()
	t.Log(f)
}

func TestDomainAnsibleBasic2(t *testing.T) {
	d := &Domain{Name: "Fool",
		Vars:   []*Var{&Var{Name: "VName", Value: "deepvalue"}, &Var{Name: "VName1", Value: "deepvalue2"}},
		Groups: []*Group{&Group{Name: "GName", Hosts: []*Host{&Host{Name: "GHName"}}}},
		Hosts: []*Host{&Host{Name: "HName"}, &Host{Name: "HName2",
			Vars: []*Var{&Var{Name: "Hello", Value: "There"}}}}}
	f := d.MarshalAnsible()
	t.Log(f)
}
