package doittypes

import (
	"database/sql"
	"time"
)

type Host struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string        `sql:"unique" json:"name"`
	Vars      []*Var        `gorm:"many2many:hostvars_vars;" json:"vars,omitempty"`
	Domain    *Domain       `json:"domain"`
	DomainID  sql.NullInt64 `json:"-"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}
