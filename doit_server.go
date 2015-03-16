package main

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
func (ds *DoitServer) Add(name string) {
	h := &Host{Name: name}
	ds.Store.Conn.NewRecord(h) // => returns `true` if primary key is blank
	ds.Store.Conn.Create(&h)

}

func (ds *DoitServer) AddHost(name string) {
}

func (ds *DoitServer) UpdateHost() {

}

func (ds *DoitServer) RemoveHost() {

}

func (ds *DoitServer) GetHost() {

}

//Group handlers
func (ds *DoitServer) AddGroup() {

}

func (ds *DoitServer) UpdateGroup() {

}

func (ds *DoitServer) RemoveGroup() {

}

func (ds *DoitServer) GetGroup() {

}

//Domain handlers
func (ds *DoitServer) AddDomain() {

}

func (ds *DoitServer) UpdateDomain() {

}

func (ds *DoitServer) RemoveDomain() {

}

func (ds *DoitServer) GetDomain() {

}

//Var handlers
func (ds *DoitServer) AddVar() {

}

func (ds *DoitServer) UpdateVar() {

}

func (ds *DoitServer) RemoveVar() {

}

func (ds *DoitServer) GetVar() {

}
