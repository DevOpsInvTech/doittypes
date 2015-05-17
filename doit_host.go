package doittypes

import (
	"database/sql"
	"reflect"
	"time"
)

//Host a representation of a host entity
type Host struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id" ansible:"-"`
	Name      string        `json:"name" ansible:"root"`
	Vars      []*Var        `gorm:"many2many:hostvars_vars;" json:"vars,omitempty" ansible:"vars"`
	Domain    *Domain       `json:"-" ansible:"-"`
	DomainID  sql.NullInt64 `json:"-" ansible:"-"`
	Group     *Group        `json:"-" ansible:"-" ansible:"parent"`
	GroupID   sql.NullInt64 `json:"-" ansible:"-"`
	CreatedAt time.Time     `json:"-" ansible:"-"`
	UpdatedAt time.Time     `json:"-" ansible:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (h *Host) MarshalAnsible(n *AnsibleNode) {
	val := reflect.ValueOf(h).Elem()
	for i := 0; i < val.NumField(); i = i + 1 {
		AnsibleCheckTag(reflect.TypeOf(h).Elem().Field(i), n, h)
	}
	for i := range h.Vars {
		h.Vars[i].MarshalAnsible(n)
	}
}
