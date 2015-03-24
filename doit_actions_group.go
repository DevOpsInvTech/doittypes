package main

//AddGroup Add group to datastore
func (ds *DoitServer) AddGroup(name string) (g *Group, err error) {
	g = &Group{Name: name}
	ds.Store.Conn.NewRecord(g)
	gormErr := ds.Store.Conn.Create(&g)
	return g, gormErr.Error
}

//AddGroupVars Add new Vars to Host
func (ds *DoitServer) AddGroupVars(id int, vars ...Var) error {
	g, err := ds.GetGroup(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Vars").Append(&vars)
	return gormErr.Error
}

//RemoveGroupVars Remove Vars from Host
func (ds *DoitServer) RemoveGroupVars(id int, vars ...Var) error {
	g, err := ds.GetGroup(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Vars").Delete(&vars)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//AddGroupHosts Add new Vars to Host
func (ds *DoitServer) AddGroupHosts(id int, hosts ...Host) error {
	g, err := ds.GetGroup(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Hosts").Append(hosts)
	return gormErr.Error
}

//RemoveGroupHosts Remove Vars from Host
func (ds *DoitServer) RemoveGroupHosts(id int, hosts ...Host) error {
	g, err := ds.GetGroup(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Hosts").Delete(&hosts)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//AddGroupDomains Add new Vars to Host
func (ds *DoitServer) AddGroupDomains(id int, domains ...Domain) error {
	g, err := ds.GetGroup(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Domains").Append(domains)
	return gormErr.Error
}

//RemoveGroupDomains Remove Vars from Host
func (ds *DoitServer) RemoveGroupDomains(id int, domains ...Domain) error {
	g, err := ds.GetGroup(id)
	if err != nil {
		return err
	}
	gormErr := ds.Store.Conn.Model(&g).Association("Domains").Delete(&domains)
	if gormErr != nil {
		return gormErr.Error
	}
	return nil
}

//RemoveGroup Remove group and its relationships to other objects
func (ds *DoitServer) RemoveGroup(id int) error {
	g, err := ds.GetGroup(id)
	if err != nil {
		return err
	}
	if len(g.Vars) > 0 {
		gormErr := ds.Store.Conn.Model(&g).Association("Vars").Delete(&g.Vars)
		if gormErr.Error != nil {
			return gormErr.Error
		}
	}
	if len(g.Hosts) > 0 {
		gormErr := ds.Store.Conn.Model(&g).Association("Hosts").Delete(&g.Hosts)
		if gormErr.Error != nil {
			return gormErr.Error
		}
	}
	hostErr := ds.Store.Conn.Delete(&g)
	if hostErr.Error != nil {
		return hostErr.Error
	}
	return nil
}

//GetGroup Get Var from datastore
func (ds *DoitServer) GetGroup(id int) (*Group, error) {
	g := &Group{ID: id}
	gormErr := ds.Store.Conn.First(&g).Related(&g.Vars, "Vars").Related(&g.Hosts, "Hosts")
	if gormErr.Error != nil {
		return g, gormErr.Error
	}
	return g, nil
}
