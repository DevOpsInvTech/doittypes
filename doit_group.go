package doittypes

import (
	"database/sql"
	"fmt"
	"time"
)

//Group a container for host and var entities
type Group struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string        `json:"name"`
	Domain    *Domain       `json:"-"`
	DomainID  sql.NullInt64 `json:"-"`
	Hosts     []*Host       `gorm:"many2many:group_hosts;" json:"hosts,omitempty"`
	Vars      []*Var        `gorm:"many2many:group_vars;" json:"vars,omitempty"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (g *Group) MarshalAnsible() map[string]interface{} {
	//vars
	var vStrs, hStrs []map[string]interface{}
	var hm, vm map[string]interface{}

	if len(g.Hosts) > 0 {
		for i := range g.Hosts {
			hStrs = append(hStrs, g.Hosts[i].MarshalAnsible())
		}
		hm = map[string]interface{}{"hosts": hStrs}
	}

	if len(g.Vars) > 0 {
		for i := range g.Vars {
			vStrs = append(vStrs, g.Vars[i].MarshalAnsible())
		}
		vm = map[string]interface{}{"vars": vStrs}
	}

	if len(g.Hosts) > 0 && len(g.Vars) > 0 {
		fmt.Println(1, 1)
		return map[string]interface{}{g.Name: []map[string]interface{}{hm, vm}}
	} else if len(g.Hosts) == 0 && len(g.Vars) > 0 {
		fmt.Println(0, 1)
		return map[string]interface{}{g.Name: vm}
	} else if len(g.Hosts) > 0 && len(g.Vars) == 0 {
		fmt.Println(1, 0)
		return map[string]interface{}{g.Name: hm}
	}
	return map[string]interface{}{g.Name: ""}
}

type GroupMatrix struct {
	ID      int `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	GroupID int
	Groups  []Group `gorm:"many2many:group_groupmatrixes;"`
}
