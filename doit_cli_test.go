package main

import (
	"os"
	"testing"
)

func TestBasicCommandLineParse(t *testing.T) {
	os.Args[1] = "-h"
	cli := &CommandLine{}
	err := cli.Usage()
	if err != nil {
		t.Log("ERR")
		t.Fatal(err)
	}
	t.Log("Fin")
}
