package doittypes

import "time"

//Domain a container for all related objects under the same namespace
type Domain struct {
	ID        int       `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string    `sql:"unique" json:"name"`
	Hosts     []*Host   `json:"hosts,omitempty"`
	Vars      []*Var    `json:"vars,omitempty"`
	Groups    []*Group  `json:"groups,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (d *Domain) MarshalAnsible() map[string]interface{} {
	var hStr []map[string]interface{}
	if len(d.Hosts) > 0 {
		for i := range d.Hosts {
			hStr = append(hStr, d.Hosts[i].MarshalAnsible())
		}
	}
	var vStr []map[string]interface{}
	if len(d.Vars) > 0 {
		for i := range d.Vars {
			vStr = append(vStr, d.Vars[i].MarshalAnsible())
		}
	}
	var gStr []map[string]interface{}
	if len(d.Groups) > 0 {
		for i := range d.Groups {
			mt := d.Groups[i].MarshalAnsible()
			for k := range mt {
				if k != "" {
					gStr = append(gStr, mt)

				}
			}
		}
	}
	return map[string]interface{}{"hosts": hStr, "vars": vStr, "groups": gStr}
}
