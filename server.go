package main

import (
	"net/http"
	"time"
)

type server struct {
	game    *Game
	port    string
	router  *http.ServeMux
	clients []*client
	server  *http.Server
}

func newServer() *server {
	return &server{
		game:   newGame(),
		port:   *flagPort,
		router: http.NewServeMux(),
		server: &http.Server{
			Addr:           *flagPort,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *server) run() error {
	s.routes()
	s.server.Handler = s.router
	return s.server.ListenAndServe()
}
