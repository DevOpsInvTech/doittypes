package main

type Domain struct {
	ID   int    `sql:"not null;unique;AUTO_INCREMENT"`
	Name string `sql:"type:varchar"`
}
