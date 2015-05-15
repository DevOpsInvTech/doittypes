package doittypes

import (
	"database/sql"
	"time"
)

type Group struct {
	ID        int           `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string        `sql:"unique" json:"name"`
	Domain    *Domain       `json:"domain"`
	DomainID  sql.NullInt64 `json:"-"`
	Hosts     []*Host       `gorm:"many2many:group_hosts;" json:"hosts,omitempty"`
	Vars      []*Var        `gorm:"many2many:group_vars;" json:"vars,omitempty"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}

type GroupMatrix struct {
	ID      int `sql:"not null;unique;AUTO_INCREMENT"`
	GroupID int
	Groups  []Group `gorm:"many2many:group_groupmatrixes;"`
}
