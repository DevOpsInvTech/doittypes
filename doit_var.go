package doittypes

import (
	"database/sql"
	"time"
)

//Var Variable to add
type Var struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Domain    *Domain       `json:"-" `
	DomainID  sql.NullInt64 `json:"-"`
	Host      *Host         `json:"-"`
	HostID    sql.NullInt64 `json:"-"`
	Group     *Group        `json:"-"`
	GroupID   sql.NullInt64 `json:"-"`
	Name      string        `json:"name"`
	Value     string        `json:"value"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (v *Var) MarshalAnsible() map[string]interface{} {
	m := map[string]interface{}{v.Name: v.Value}
	return m
}
