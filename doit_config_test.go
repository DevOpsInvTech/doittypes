package main

import "testing"

func TestLoadConfig(t *testing.T) {
	dc := &DoitConfig{}
	dc.Read("test_configs/test-config.yml")
}
