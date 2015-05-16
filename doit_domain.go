package doittypes

import "time"

//Domain a container for all related objects under the same namespace
type Domain struct {
	ID        int       `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string    `sql:"unique" json:"name"`
	Hosts     []Host    `gorm:"many2many:domain_hosts;" json:"hosts"`
	Vars      []Var     `gorm:"many2many:domain_vars;"json:"vars"`
	Groups    []Group   `gorm:"many2many:domain_groups;"json:"groups"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
