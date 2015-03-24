package main

import (
	"flag"
	"log"
)

func main() {
	dc := &DoitConfig{}

	port := flag.String("p", "8080", "Port")
	config := flag.String("c", "", "Load config file")
	serverMode := flag.Bool("s", false, "Enable server mode")

	flag.Parse()

	log.Println(port, serverMode, config)

	if *config != "" {
		//load config
		err := dc.Read(*config)
		if err != nil {
			panic(err)
		}
	} else {
		//manual load config

	}

	if *serverMode {
		ds := &DoitServer{Store: &DoitStorage{}}
		ds.Listen(port, dc)
	} else {
		//TODO: Act as a CLI client
	}
}
