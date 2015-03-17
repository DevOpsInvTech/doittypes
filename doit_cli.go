package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

type CommandLine struct {
}

func (cli *CommandLine) Usage() {
	usage := `doit-client

	Usage:
	  doit-client add host <name>
	  doit-client remove host <name>
	  doit-client get host <name>
	`
	_, err := docopt.Parse(usage, nil, true, "DOIT Client", false)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(args)

}
