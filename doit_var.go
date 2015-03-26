package main

import (
	"database/sql"
	"time"
)

type Var struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Domain    *Domain       `json:"domain"`
	DomainID  sql.NullInt64 `json:"-"`
	Name      string        `sql:"unique json:"name""`
	Value     string        `json:"value"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}
