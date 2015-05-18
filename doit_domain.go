package doittypes

import (
	"fmt"
	"strings"
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
func (d *Domain) MarshalAnsible() string {
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
	var hStr []string
	if len(d.Hosts) > 0 {
		for i := range d.Hosts {
			hStr = append(hStr, d.Hosts[i].MarshalAnsible())
		}
	}
	var vStr []string
	if len(d.Vars) > 0 {
		for i := range d.Vars {
			vStr = append(vStr, d.Vars[i].MarshalAnsible())
		}
	}
	var gStr []string
	if len(d.Groups) > 0 {
		for i := range d.Groups {
			gStr = append(gStr, d.Groups[i].MarshalAnsible())
		}
	}
	var a, a1, b, b1, c, c1 string
	if len(hStr) > 0 {
		a = strings.Join(hStr, ",")
		if len(hStr) == 1 {
			a1 = fmt.Sprintf("\"hosts\":%s", a)
		} else if len(vStr) > 1 {
			a1 = fmt.Sprintf("\"hosts\":[%s]", a)
		}
	}
	if len(vStr) > 0 {
		b = strings.Join(vStr, ",")
		if len(vStr) == 1 {
			b1 = fmt.Sprintf("\"vars\":%s", b)
		} else if len(vStr) > 1 {
			b1 = fmt.Sprintf("\"vars\":[%s]", b)
		}
	}
	if len(gStr) > 0 {
		c = strings.Join(gStr, ",")
		if len(gStr) == 1 {
			c1 = fmt.Sprintf("\"groups\":%s", c)
		} else if len(vStr) > 1 {
			c1 = fmt.Sprintf("\"groups\":[%s]", c)
		}
	}
	return fmt.Sprintf("{%s}", strings.Join([]string{a1, b1, c1}, ","))
}
