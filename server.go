package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type server struct {
	game    *Game
	port    string
	router  *http.ServeMux
	clients []*client
	msg     chan []byte
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

func (s *server) sendState() {

	for _, c := range s.clients {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetIndent("", "  ")
		s.game.State()
		ps, _ := s.playerState(c.playerID)
		err := enc.Encode(ps)
		if err != nil {
			log.Println("sendState encoding error:", err)
		}
		b := buf.Bytes()
		c.messages <- b
	}
}

type PlayerState struct {
	Player   Player   `json:"player"`
	Oponents []Player `json:"oponent"`
	HeapHead Card     `json:"heap_head"`
}

func (s *server) playerState(id string) (PlayerState, bool) {
	ps := PlayerState{}
	found := false
	s.game.State()
	for _, p := range s.game.Players {
		switch p.ID {
		case id:
			ps.Player = *p
			found = true
		default:
			// Hide oponent cards and id
			oponent := Player{
				ID:   "",
				Name: p.Name,
				Cards: &CardStack{
					Cards: make([]Card, len(p.Cards.Cards)),
				},
			}
			for i := range p.Cards.Cards {
				oponent.Cards.Cards[i].Color = "back"
			}
			ps.Oponents = append(ps.Oponents, oponent)
		}
	}
	ps.HeapHead = s.game.HeapHead
	// Just show the label for the HeapHead
	// so  9 - 3 is shown instead of the value
	ps.HeapHead.Value = ps.HeapHead.Label
	return ps, found
}

func (s *server) runGame() {
	for {

	}
}
