package main

import "errors"

//Domain handlers
func (ds *DoitServer) AddDomain(name string) {
	d := &Domain{Name: name}
	ds.Store.Conn.NewRecord(d)
	ds.Store.Conn.Create(&d)
}

func (ds *DoitServer) UpdateDomain() {

}

func (ds *DoitServer) RemoveDomain(id int) error {
	d := &Domain{ID: id}
	ds.Store.Conn.First(&d)
	if d.Name != "" {
		ds.Store.Conn.Delete(d)
		return nil
	}
	return errors.New("Domain ID not found")
}

func (ds *DoitServer) GetDomain(id int) (*Domain, error) {
	d := &Domain{ID: id}
	ds.Store.Conn.First(&d)
	if d.Name != "" {
		return d, nil
	}
	return nil, errors.New("Domain ID not found")
}
