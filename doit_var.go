package doittypes

import (
	"database/sql"
	"log"
	"reflect"
	"time"
)

//Var Variable to add
type Var struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id" ansible:"-"`
	Domain    *Domain       `json:"-" ansible:"-" `
	DomainID  sql.NullInt64 `json:"-" ansible:"-"`
	Host      *Host         `json:"-" ansible:"parent"`
	HostID    sql.NullInt64 `json:"-" ansible:"-"`
	Group     *Group        `json:"-" ansible:"parent"`
	GroupID   sql.NullInt64 `json:"-" ansible:"-"`
	Name      string        `json:"name" ansible:"root"`
	Value     string        `json:"value" ansible:"value"`
	CreatedAt time.Time     `json:"-" ansible:"-"`
	UpdatedAt time.Time     `json:"-" ansible:"-"`
}

//MarshalAnsilbe mashals the struct into an Ansible supported JSON
func (v *Var) MarshalAnsilbe() {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i = i + 1 {
		field := reflect.TypeOf(v).Elem().Field(i)
		log.Println(field.Tag.Get(AnsileTag))
	}
}
