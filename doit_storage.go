package main

import "github.com/jinzhu/gorm"

type DoitStorage struct {
	Type     string
	Location string
	Conn     gorm.DB
}

func NewStorage(t string, loc string) (*DoitStorage, error) {
	db, err := gorm.Open(t, loc)
	if err != nil {
		return nil, err
	}

	s := &DoitStorage{Conn: db, Type: t, Location: loc}
	s.Conn.DB()
	db.DB().Ping()
	return s, nil
}

func (s *DoitStorage) InitSchema(overwrite bool) {
	if overwrite {
		s.Conn.CreateTable(&Host{})
		s.Conn.CreateTable(&HostVar{})
		s.Conn.CreateTable(&Var{})
		s.Conn.CreateTable(&Domain{})
		s.Conn.CreateTable(&Group{})
		s.Conn.CreateTable(&GroupMatrix{})
	} else {
		//TODO: Test schema
		//test schema
	}

}

func (s *DoitStorage) Close() error {
	err := s.Conn.Close()
	return err
}
