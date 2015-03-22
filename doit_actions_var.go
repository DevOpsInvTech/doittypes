package main

import "errors"

//Var handlers
func (ds *DoitServer) AddVar(name string, value string) {
	v := &Var{Name: name, Value: value}
	ds.Store.Conn.NewRecord(v)
	ds.Store.Conn.Create(&v)
}

func (ds *DoitServer) UpdateVar(id int, value string) error {
	v := &Var{ID: id}
	ds.Store.Conn.First(&v)
	if v.Name != "" {
		v.Value = value
		ds.Store.Conn.Save(&v)
		return nil
	}
	return errors.New("Var ID not found")
}

func (ds *DoitServer) RemoveVar(id int) error {
	v := &Var{ID: id}
	ds.Store.Conn.First(&v)
	if v.Name != "" {
		ds.Store.Conn.Delete(&v)
		return nil
	}
	return errors.New("Var ID not found")
}

func (ds *DoitServer) GetVar(id int) (*Var, error) {
	v := &Var{ID: id}
	ds.Store.Conn.First(&v)
	if v.Name != "" {
		return v, nil
	}
	return nil, errors.New("Var ID not found")
}
