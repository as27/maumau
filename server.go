package main

import "net/http"

type server struct {
	game    *Game
	port    string
	router  *http.ServeMux
	clients []*client
}

func newServer() *server {
	return &server{
		game:   newGame(),
		port:   *flagPort,
		router: http.NewServeMux(),
	}
}

func (s *server) run() error {
	return nil
}
