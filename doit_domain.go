package doittypes

import (
	"log"
	"time"
)

//Domain a container for all related objects under the same namespace
type Domain struct {
	ID        int       `sql:"not null;unique;AUTO_INCREMENT" json:"id" ansible:"-"`
	Name      string    `sql:"unique" json:"name" ansible:"-"`
	Hosts     []*Host   `json:"hosts,omitempty" ansible:"main"`
	Vars      []*Var    `json:"vars,omitempty" ansible:"main"`
	Groups    []*Group  `json:"groups,omitempty" ansible:"main"`
	CreatedAt time.Time `json:"-" ansible:"-"`
	UpdatedAt time.Time `json:"-" ansible:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (d *Domain) MarshalAnsible(n *AnsibleNode) {
	//Check every sub struct
	//Marshal each sub struct
	//Hosts, Vars, Groups
	// check for ansible tag
	// - means skip
	// root means take that as the root value
	// parent means to link the item to its parent
	// value means value for root item
	// main means this is a top tier item to embed into the json
	//dJSON := make(map[string]interface{})
	n := &AnsibleNode{}
	for i := range d.Hosts {
		d.Hosts[i].MarshalAnsible(n)
	}
	for i := range d.Vars {
		d.Vars[i].MarshalAnsible(n)
	}
	for i := range d.Groups {
		d.Groups[i].MarshalAnsible(n)
	}
	log.Printf("%#v", n)
}
