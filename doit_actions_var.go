package main

//AddVar Add new var to datastore
func (ds *DoitServer) AddVar(d *Domain, name string, value string) (v *Var, err error) {
	v = &Var{Name: name, Value: value, Domain: d}
	ds.Store.Conn.NewRecord(v)
	gormErr := ds.Store.Conn.Create(&v)
	return v, gormErr.Error
}

//UpdateVar Update Var
func (ds *DoitServer) UpdateVar(d *Domain, id int, value string) error {
	v, err := ds.GetVar(d, id)
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
func (ds *DoitServer) RemoveVar(d *Domain, id int) error {
	v, err := ds.GetVar(d, id)
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
func (ds *DoitServer) GetVar(d *Domain, id int) (*Var, error) {
	v := &Var{ID: id, Domain: d}
	gormErr := ds.Store.Conn.First(&v)
	if gormErr.Error != nil {
		return v, gormErr.Error
	}
	return v, nil
}

//GetVarByName Get Var from datastore
func (ds *DoitServer) GetVarByName(d *Domain, name string) (*Var, error) {
	v := &Var{Name: name, Domain: d}
	gormErr := ds.Store.Conn.First(&v)
	if gormErr.Error != nil {
		return v, gormErr.Error
	}
	return v, nil
}
