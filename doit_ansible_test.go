package doittypes

import "testing"

func TestDomainAnsibleBasic(t *testing.T) {
	d := &Domain{Name: "Fool",
		Vars:   []*Var{&Var{Name: "deepvar", Value: "deepvalue"}},
		Groups: []*Group{&Group{Name: "Group"}},
		Hosts: []*Host{&Host{Name: "Host1",
			Vars: []*Var{&Var{Name: "Hello", Value: "There"}}}}}
	d.MarshalAnsilbe()
}
