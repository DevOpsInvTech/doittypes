package main

import "time"

type Var struct {
	ID        int `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
