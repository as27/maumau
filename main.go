package main

import (
	"flag"
	"log"
)

var (
	flagPort = flag.String("-p", ":2704", "Port für den Server")
)

func main() {
	flag.Parse()
	s := newServer()
	s.routes()
	err := s.run()
	if err != nil {
		log.Fatal("error main.s.run():", err)
	}
}
