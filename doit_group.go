package doittypes

import (
	"database/sql"
	"log"
	"reflect"
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
func (g *Group) MarshalAnsilbe() {
	val := reflect.ValueOf(g).Elem()
	for i := 0; i < val.NumField(); i = i + 1 {
		field := reflect.TypeOf(g).Elem().Field(i)
		log.Println(field.Tag.Get(AnsileTag))
	}
	for i := range g.Hosts {
		g.Hosts[i].MarshalAnsilbe()
	}
	for i := range g.Vars {
		g.Vars[i].MarshalAnsilbe()
	}
}

type GroupMatrix struct {
	ID      int `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	GroupID int
	Groups  []Group `gorm:"many2many:group_groupmatrixes;"`
}
