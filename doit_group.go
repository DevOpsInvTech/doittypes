package main

import "time"

type Group struct {
	ID        int      `sql:"not null;unique;AUTO_INCREMENT"`
	Name      string   `sql:"unique"`
	Hosts     []Host   `gorm:"many2many:group_hosts;"`
	Vars      []Var    `gorm:"many2many:group_vars;"`
	Domains   []Domain `gorm:"many2many:group_domains;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GroupMatrix struct {
	ID      int `sql:"not null;unique;AUTO_INCREMENT"`
	GroupID int
	Groups  []Group `gorm:"many2many:group_groupmatrixes;"`
}

func (gm *GroupMatrix) AddGroup(g Group) {
	gm.Groups = append(gm.Groups, g)
}
