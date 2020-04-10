package main

import (
	"flag"
	"log"
)

var (
	flagPort = flag.String("p", ":2704", "Port für den Server")
	flagCalc = flag.Int("c", 0, "Erstellt für die Ablage eine Rechenaufgabe mit dem Zahlenraum der über dieses Flag definiert wird.")
)

func main() {
	flag.Parse()
	s := newServer()
	log.Println("Starting server at ", *flagPort)
	err := s.run()
	if err != nil {
		log.Fatal("error main.s.run():", err)
	}
}
