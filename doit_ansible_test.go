package doittypes

import "testing"

func TestDomainAnsibleBasic(t *testing.T) {
	d := &Domain{Name: "Fool",
		Vars:   []*Var{&Var{Name: "VName", Value: "deepvalue"}},
		Groups: []*Group{&Group{Name: "GName"}},
		Hosts: []*Host{&Host{Name: "HName",
			Vars: []*Var{&Var{Name: "Hello", Value: "There"}}}}}
	d.MarshalAnsible()
}
