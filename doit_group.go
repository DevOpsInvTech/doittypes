package doittypes

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Group a container for host and var entities
type Group struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id" ansible:"-"`
	Name      string        `json:"name" ansible:"root"`
	Domain    *Domain       `json:"-" ansible:"-"`
	DomainID  sql.NullInt64 `json:"-" ansible:"-"`
	Hosts     []*Host       `gorm:"many2many:group_hosts;" json:"hosts,omitempty" ansible:"hosts"`
	Vars      []*Var        `gorm:"many2many:group_vars;" json:"vars,omitempty" ansible:"vars"`
	CreatedAt time.Time     `json:"-" ansible:"-"`
	UpdatedAt time.Time     `json:"-" ansible:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (g *Group) MarshalAnsible() string {
	//vars
	var vStrs, hStrs []string
	if len(g.Hosts) > 0 {
		for i := range g.Hosts {
			hStrs = append(hStrs, g.Hosts[i].MarshalAnsible())
		}
	}
	if len(g.Vars) > 0 {
		for i := range g.Vars {
			vStrs = append(vStrs, g.Vars[i].MarshalAnsible())
		}
	}
	var a, a1, b, b1 string
	if len(hStrs) > 0 {
		a = strings.Join(hStrs, ",")
		if len(hStrs) == 1 {
			a1 = fmt.Sprintf("\"hosts\":%s", a)
		} else if len(vStrs) > 1 {
			a1 = fmt.Sprintf("\"hosts\":[%s]", a)
		}
	}
	if len(vStrs) > 0 {
		b = strings.Join(vStrs, ",")
		if len(vStrs) == 1 {
			b1 = fmt.Sprintf("\"vars\":%s", b)
		} else if len(vStrs) > 1 {
			b1 = fmt.Sprintf("\"vars\":[%s]", b)
		}
	}
	if len(vStrs) > 0 && len(hStrs) > 0 {
		return fmt.Sprintf("\"%s\":%s", g.Name, strings.Join([]string{a1, b1}, ","))
	} else if len(vStrs) == 0 && len(hStrs) == 1 {
		return fmt.Sprintf("\"%s\":%s", g.Name, a1)
	} else if len(vStrs) == 1 && len(hStrs) == 0 {
		return fmt.Sprintf("\"%s\":%s", g.Name, b1)
	} else if len(vStrs) == 0 && len(hStrs) == 0 {
		return fmt.Sprintf("\"%s\"", g.Name)
	}
	return fmt.Sprintf("\"%s\"", g.Name)
}

type GroupMatrix struct {
	ID      int `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	GroupID int
	Groups  []Group `gorm:"many2many:group_groupmatrixes;"`
}
