package main

type Host struct {
	ID   int
	Name string
	Vars []*Var
}

func (h *Host) AddVar(v *Var) {
	h.Vars = append(h.Vars, v)
}
