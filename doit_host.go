package doittypes

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//Host a representation of a host entity
type Host struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string        `json:"name"`
	Vars      []*Var        `gorm:"many2many:hostvars_vars;" json:"vars,omitempty"`
	Domain    *Domain       `json:"-"`
	DomainID  sql.NullInt64 `json:"-"`
	Group     *Group        `json:"-"`
	GroupID   sql.NullInt64 `json:"-"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (h *Host) MarshalAnsible() string {
	var vStr []string
	if len(h.Vars) > 0 {
		for i := range h.Vars {
			vStr = append(vStr, h.Vars[i].MarshalAnsible())
		}
		if len(h.Vars) == 1 {
			return fmt.Sprintf("\"%s\":{\"vars\":%s}", h.Name, strings.Join(vStr, ","))
		} else if len(h.Vars) > 1 {
			return fmt.Sprintf("{\"%s\":{\"vars\":[%s]}}", h.Name, strings.Join(vStr, ","))
		}
	}
	return fmt.Sprintf("\"%s\"", h.Name)
}
