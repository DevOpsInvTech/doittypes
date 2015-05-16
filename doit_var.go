package doittypes

import (
	"database/sql"
	"time"
)

type Var struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Domain    *Domain       `json:"-"`
	DomainID  sql.NullInt64 `json:"-"`
	Host      *Host         `json:"-"`
	HostID    sql.NullInt64 `json:"-"`
	Group     *Group        `json:"-"`
	GroupID   sql.NullInt64 `json:"-"`
	Name      string        `json:"name""`
	Value     string        `json:"value"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}
