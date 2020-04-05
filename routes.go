package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *server) routes() {
	s.router.HandleFunc("/gamestate", s.handleGameState())
	s.router.HandleFunc("/ws", s.handleWS())
	s.router.HandleFunc("/start", s.handleStart())
	s.router.HandleFunc("/", s.handleRoot())
}

func (s *server) handleWS() http.HandlerFunc {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		player := newPlayer("")
		s.game.Event(addPlayer(player))
		log.Println("New Player:", player)
		c := &client{
			socket:   conn,
			messages: make(chan []byte, 256), // message buffer 256 bytes
			playerID: player.ID,
		}
		s.clients = append(s.clients, c)
	}
}
func (s *server) handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/gametable.html")
	}
}

func (s *server) handleGameState() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		s.game.State()
		err := enc.Encode(s.game)
		if err != nil {
			log.Println("handleGameState error:", err)
		}
	}
}

func (s *server) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cardGame := CardGame()
		cardGame.shuffle()
		s.game.Event(addCardGameToStack(cardGame))
		s.game.Event(serveGame())
	}
}
