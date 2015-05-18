package doittypes

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDomainAnsibleBasic1(t *testing.T) {
	d := &Domain{Name: "Fool",
		Vars:   []*Var{&Var{Name: "VName", Value: "deepvalue"}},
		Groups: []*Group{&Group{Name: "GName"}},
		Hosts: []*Host{&Host{Name: "HName",
			Vars: []*Var{&Var{Name: "Hello", Value: "There"}}}}}
	f := d.MarshalAnsible()
	t.Logf("%#v", f)
	j, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(j))
}

func TestDomainAnsibleBasic2(t *testing.T) {
	d := &Domain{Name: "Fool",
		Vars:   []*Var{&Var{Name: "VName", Value: "deepvalue"}, &Var{Name: "VName1", Value: "deepvalue2"}},
		Groups: []*Group{&Group{Name: "GName", Hosts: []*Host{&Host{Name: "GHName"}}}},
		Hosts: []*Host{&Host{Name: "HName"}, &Host{Name: "HName2",
			Vars: []*Var{&Var{Name: "Hello", Value: "There"}}}}}
	f := d.MarshalAnsible()
	t.Logf("%#v", f)
	j, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(j))
}

func TestDomainAnsibleBasic3(t *testing.T) {
	d := &Domain{Name: "Fool",
		Vars: []*Var{&Var{Name: "VName", Value: "deepvalue"},
			&Var{Name: "VName1", Value: "deepvalue2"}},
		Groups: []*Group{&Group{Name: "GName",
			Hosts: []*Host{&Host{Name: "GHName"}}}},
		Hosts: []*Host{&Host{Name: "HName",
			Vars: []*Var{&Var{Name: "aaaa", Value: "bbbb"}}},
			&Host{Name: "HName2",
				Vars: []*Var{&Var{Name: "Hello", Value: "There"}}}}}
	f := d.MarshalAnsible()
	t.Logf("%#v", f)
	j, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(j))
}

func TestDomainAnsibleGroup1(t *testing.T) {
	d := &Domain{Name: "Fool",
		Groups: []*Group{&Group{Name: "GName1"}, &Group{Name: "GName2"}, &Group{Name: "GName3"}},
	}
	f := d.MarshalAnsible()
	t.Logf("%#v", f)
	j, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(j))
}
