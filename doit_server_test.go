package main

import "testing"

func TestDoitServerAddHost(t *testing.T) {
	ds := &DoitServer{}
	ds.OpenDatastore("sqlite3", "_test_tmp/TestDoitServerAddHost.db")
	ds.Store.InitSchema(true)
	ds.AddHost("Steve")
	ds.CloseDatastore()
}
