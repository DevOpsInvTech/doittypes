package main

import (
	"errors"
	"fmt"
)

//AddHost Add new host to the datastore
func (ds *DoitServer) AddHost(d *Domain, name string) (h *Host, err error) {
	domain, err := ds.GetDomain(d.ID)
	if err != nil {
		return h, err
	}
	h = &Host{Name: name, Domain: domain}
	ds.Store.Conn.NewRecord(h)
	gormErr := ds.Store.Conn.Create(&h)
	return h, gormErr.Error
}

//AddHostVars Add new Vars to Host
func (ds *DoitServer) AddHostVars(d *Domain, id int, vars ...HostVar) error {
	h, err := ds.GetHost(d, id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&h).Association("Vars").Append(vars)
	return gormErr.Error
}

//RemoveHostVars Remove Vars from Host
func (ds *DoitServer) RemoveHostVars(d *Domain, id int, vars ...HostVar) error {
	h, err := ds.GetHost(d, id)
	if err != nil {
		return err
	}
	for i, v := range vars {
		fmt.Println(i)
		rmVar, err := ds.GetHostVar(d, v.ID)
		fmt.Println(rmVar)
		if err != nil {
			return err
		}
		varErr := ds.Store.Conn.Delete(&rmVar)
		if varErr.Error != nil {
			return varErr.Error
		}
	}
	gormErr := ds.Store.Conn.Model(&h).Association("Vars").Delete(&vars)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//RemoveHost Remove host from datastore
func (ds *DoitServer) RemoveHost(d *Domain, id int) error {
	h, err := ds.GetHost(d, id)
	if err != nil {
		return err
	}
	if len(h.Vars) > 0 {
		gormErr := ds.Store.Conn.Model(&h).Association("Vars").Delete(&h.Vars)
		if gormErr.Error != nil {
			return gormErr.Error
		}
	}
	hostErr := ds.Store.Conn.Delete(&h)
	if hostErr.Error != nil {
		return hostErr.Error
	}
	return nil
}

//GetHost Get host from datastore
func (ds *DoitServer) GetHost(d *Domain, id int) (*Host, error) {
	h := &Host{ID: id, Domain: d}
	gormErr := ds.Store.Conn.First(&h).Related(&h.Vars, "Vars")
	if gormErr.Error != nil {
		return h, gormErr.Error
	}
	return h, nil
}

//GetHostVar Get HostVar from datastore
func (ds *DoitServer) GetHostVar(d *Domain, id int) (*HostVar, error) {
	v := &HostVar{ID: id, Domain: d}
	ds.Store.Conn.First(&v)
	if v.Name != "" {
		return v, nil
	}
	return nil, errors.New("HostVar ID not found")
}
