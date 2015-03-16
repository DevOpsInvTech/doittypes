package main

import "github.com/jinzhu/gorm"

type DoitStorage struct {
	Type     string
	Location string
	DB       gorm.DB
}

func NewStorage(t string, loc string) (*DoitStorage, error) {
	db, err := gorm.Open(t, loc)
	if err != nil {
		return nil, err
	}

	return &DoitStorage{DB: db, Type: t, Location: loc}, nil
}

func (s *DoitStorage) Close() error {
	err := s.DB.Close()
	return err
}
