package main

import "errors"

type DoitServer struct {
	Store *DoitStorage
}

//OpenDatastore open datastore for writing
func (ds *DoitServer) OpenDatastore(t string, loc string) error {
	s, err := NewStorage(t, loc)
	ds.Store = s
	return err
}

//CloseDatastore close datastore
func (ds *DoitServer) CloseDatastore() error {
	err := ds.Store.Close()
	return err
}

//Host handlers
func (ds *DoitServer) AddHost(name string) {
	h := &Host{Name: name}
	ds.Store.Conn.NewRecord(h) // => returns `true` if primary key is blank
	ds.Store.Conn.Create(&h)
}

func (ds *DoitServer) AddHostVars(id int, vars ...Var) error {
	h := &Host{ID: id}
	ds.Store.Conn.First(&h)
	if h.Name != "" {
		for item := range vars {
			h.AddVar(vars[item])
		}
		ds.Store.Conn.Save(&h)
		return nil
	}
	return errors.New("Host ID not found")
}

func (ds *DoitServer) RemoveHostVars(id int, vars ...Var) error {
	return errors.New("Host ID not found")
}

func (ds *DoitServer) RemoveHost(id int) error {
	h := &Host{ID: id}
	ds.Store.Conn.First(&h)
	if h.Name != "" {
		ds.Store.Conn.Delete(&h)
		return nil
	}
	return errors.New("Host ID not found")
}

func (ds *DoitServer) GetHost(id int) (*Host, error) {
	h := &Host{ID: id}
	ds.Store.Conn.First(&h)
	if h.Name != "" {
		return h, nil
	}
	return nil, errors.New("Host ID not found")
}

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
