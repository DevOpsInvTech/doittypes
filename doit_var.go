package main

import (
	"database/sql"
	"time"
)

type Var struct {
	ID        int `sql:"not null;unique;AUTO_INCREMENT"`
	Domain    *Domain
	DomainID  sql.NullInt64
	Name      string `sql:"unique"`
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
