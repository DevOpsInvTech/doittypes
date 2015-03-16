package main

import "time"

type Group struct {
	ID    int `sql:"not null;unique;AUTO_INCREMENT"`
	Name  string
	Hosts []Host `gorm:"many2many:group_hosts;"`
	Vars  []Var  `gorm:"many2many:group_vars;"`
	//Groups    []int    `gorm:"many2many:group_groups;"`
	Domain    []Domain `gorm:"many2many:group_domains;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (g *Group) AddVar(v Var) {
	g.Vars = append(g.Vars, v)
}

//func (g *Group) AddChild(c Group) {
//	g.Groups = append(g.Groups, c)
//}

func (g *Group) AddHost(h Host) {
	g.Hosts = append(g.Hosts, h)
}

type Groups struct {
	GroupID  int
	GroupID2 int
}
