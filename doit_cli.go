package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

type CommandLine struct {
}

func (cli *CommandLine) Usage() error {

	usage := `	Doit Client - DevOps Inventory Technician

	Usage:
		doit-client --version
		doit-client add host <name>
		doit-client remove host <name>
		doit-client add group <name>
		doit-client remove group <name>
		doit-client add var <name> <value>
		doit-client remove var <name>
		doit-client add domain
		doit-client remove domain

	Options:
		-h --help     			Show usage.
		--port=<port> 			Set port for server [default: 8123].
		--config=<config_file>  Load config file for server.
		-s --server      		Enable server mode.`

	arguments, err := docopt.Parse(usage, nil, true, "DOIT Client 0.1", false)
	if err != nil {
		return err
	}
	fmt.Printf("%#v", arguments)
	return nil
}
