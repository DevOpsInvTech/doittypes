package main

import "time"

type Host struct {
	ID        int `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string
	Vars      []Var
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (h *Host) AddVar(v Var) {
	h.Vars = append(h.Vars, v)
}
