package main

//AddVar Add new var to datastore
func (ds *DoitServer) AddVar(name string, value string) (v *Var, err error) {
	v = &Var{Name: name, Value: value}
	ds.Store.Conn.NewRecord(v)
	gormErr := ds.Store.Conn.Create(&v)
	return v, gormErr.Error
}

//UpdateVar Update Var
func (ds *DoitServer) UpdateVar(id int, value string) error {
	v, err := ds.GetVar(id)
	if err != nil {
		return err
	}
	v.Value = value
	gormErr := ds.Store.Conn.Save(&v)
	if gormErr.Error != nil {
		return gormErr.Error
	}
	return nil
}

//RemoveVar Remove Var
func (ds *DoitServer) RemoveVar(id int) error {
	v, err := ds.GetVar(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Delete(&v)
	if gormErr.Error != nil {
		return gormErr.Error
	}
	return nil
}

//GetVar Get Var from datastore
func (ds *DoitServer) GetVar(id int) (*Var, error) {
	v := &Var{ID: id}
	gormErr := ds.Store.Conn.First(&v)
	if gormErr.Error != nil {
		return v, gormErr.Error
	}
	return v, nil
}
