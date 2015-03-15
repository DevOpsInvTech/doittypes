package main

type Group struct {
	ID       int
	Name     string
	Hosts    []*Host
	Vars     []*Var
	Children []string
	Domain   []Domain
}

func (g *Group) AddVar(v *Var) {
	g.Vars = append(g.Vars, v)
}

func (g *Group) AddChild(c string) {
	g.Children = append(g.Children, c)
}

func (g *Group) AddHost(h *Host) {
	g.Hosts = append(g.Hosts, h)
}
