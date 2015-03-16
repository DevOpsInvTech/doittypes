package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DoitConfig struct {
	Server  ServerConfig  `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
}

type ServerConfig struct {
	Enable bool `yaml:"enable"`
	Port   int  `yaml:"port"`
}

type StorageConfig struct {
	Type     string `yaml:"type"`
	Location string `yaml:"location"`
}

func (dc *DoitConfig) Read(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	log.Println(string(data))
	nc := &DoitConfig{}
	err = yaml.Unmarshal(data, &nc)
	if err != nil {
		return err
	}
	log.Println(nc.Server.Enable)
	return nil
}
