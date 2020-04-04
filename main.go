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
	s := &server{
		game: newGame(),
		port: *flagPort,
	}
	err := s.run()
	if err != nil {
		log.Fatal("error main.s.run():", err)
	}
}
