package main

import "testing"

func TestNewStorage(t *testing.T) {
	s, err := NewStorage("sqlite3", "_test_tmp/TestNewStorage.db")
	if err != nil {
		t.Fatal(err)
	}
	s.Close()
}

func TestStorageInitSchema(t *testing.T) {
	s, err := NewStorage("sqlite3", "_test_tmp/TestStorageInitSchema.db")
	if err != nil {
		t.Fatal(err)
	}
	s.InitSchema(true)
	s.Close()
}
