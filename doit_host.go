package main

import "time"

type Host struct {
	ID        int `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string
	Vars      []Var `gorm:"many2many:host_vars;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (h *Host) AddVar(v Var) {
	h.Vars = append(h.Vars, v)
}
