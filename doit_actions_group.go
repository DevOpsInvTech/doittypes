package main

import "errors"

//Group handlers
func (ds *DoitServer) AddGroup(name string) {
	g := &Group{Name: name}
	ds.Store.Conn.NewRecord(g) // => returns `true` if primary key is blank
	ds.Store.Conn.Create(&g)
}

func (ds *DoitServer) UpdateGroup() {

}

func (ds *DoitServer) RemoveGroup(id int) error {
	g := &Group{ID: id}
	ds.Store.Conn.First(&g)
	if g.Name != "" {
		ds.Store.Conn.Delete(g)
		return nil
	}
	return errors.New("Host ID not found")
}

func (ds *DoitServer) GetGroup(id int) (*Group, error) {
	g := &Group{ID: id}
	ds.Store.Conn.First(&g)
	if g.Name != "" {
		return g, nil
	}
	return nil, errors.New("Group ID not found")
}
