package main

import "time"

type Host struct {
	ID        int       `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string    `sql:"unique"`
	Vars      []HostVar `gorm:"many2many:hostvars_vars;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HostVar struct {
	ID        int    `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string `sql:"unique"`
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
