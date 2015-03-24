package main

import (
	"database/sql"
	"time"
)

type Host struct {
	ID        int       `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string    `sql:"unique"`
	Vars      []HostVar `gorm:"many2many:hostvars_vars;"`
	Domain    Domain
	DomainID  sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HostVar struct {
	ID        int    `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string `sql:"unique"`
	Value     string
	Domain    Domain
	DomainID  sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
}
