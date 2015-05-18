package doittypes

import (
	"database/sql"
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
func (h *Host) MarshalAnsible() map[string]interface{} {
	var vStr []map[string]interface{}
	if len(h.Vars) > 0 {
		for i := range h.Vars {
			vStr = append(vStr, h.Vars[i].MarshalAnsible())
		}
		return map[string]interface{}{h.Name: map[string]interface{}{"vars": vStr}}
	}
	return map[string]interface{}{h.Name: ""}
}
