package main

import "time"

type Group struct {
	ID        int `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string
	Hosts     []Host   `gorm:"many2many:group_hosts;"`
	Vars      []Var    `gorm:"many2many:group_vars;"`
	Domains   []Domain `gorm:"many2many:group_domains;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g *Group) AddVar(v Var) {
	g.Vars = append(g.Vars, v)
}

func (g *Group) AddHost(h Host) {
	g.Hosts = append(g.Hosts, h)
}

func (g *Group) AddDomain(d Domain) {
	g.Domains = append(g.Domains, d)
}

type GroupMatrix struct {
	ID      int `sql:"not null;unique;AUTO_INCREMENT"`
	GroupID int
	Groups  []Group `gorm:"many2many:group_groupmatrixes;"`
}

func (gm *GroupMatrix) AddGroup(g Group) {
	gm.Groups = append(gm.Groups, g)
}
