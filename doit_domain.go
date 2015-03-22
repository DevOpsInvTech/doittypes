package main

import "time"

type Domain struct {
	ID        int     `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string  `sql:"type:varchar"`
	Hosts     []Host  `gorm:"many2many:domain_hosts;"`
	Vars      []Var   `gorm:"many2many:domain_vars;"`
	Groups    []Group `gorm:"many2many:domain_groups;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *Domain) AddVar(v Var) {
	d.Vars = append(d.Vars, v)
}

func (d *Domain) AddHost(h Host) {
	d.Hosts = append(d.Hosts, h)
}

func (d *Domain) AddGroup(g Group) {
	d.Groups = append(d.Groups, g)
}
