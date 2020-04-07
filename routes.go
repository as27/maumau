package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/websocket"
)

/*
Status X00 - X49 are success messages
Error status are X50 - X99
1XX -> card specific status
2XX -> player specific status
*/
const (
	StatusCardPlayed     = "100"
	StatusCardNotFound   = "150"
	StatusPlayerFound    = "200"
	StatusPlayerNotFound = "250"
)

func (s *server) routes() {
	s.router.HandleFunc("/gamestate", s.handleGameState())
	s.router.HandleFunc("/data", s.handleData())
	s.router.HandleFunc("/ws", s.handleWS())
	s.router.HandleFunc("/start", s.handleStart())
	s.router.HandleFunc("/playcard", s.handlePlayCard())
	s.router.HandleFunc("/takecard", s.handleTakeCard())
	s.router.HandleFunc("/newgame", s.handleNewGame())
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
		s.game.State()
		b := httptest.NewRecorder()
		id, ok := s.handGetID(b, r)
		if !ok {
			// Don't add a third player
			if len(s.game.Players) <= 2 {
				player := newPlayer("")
				s.game.Event(addPlayer(player))
				id = player.ID
				log.Println("New Player:", player)
			}
		}
		// allow more client instances
		// when there is a reload at the browser the id of the
		// url is used to identify the user.
		c := &client{
			socket:   conn,
			messages: make(chan []byte, 256), // message buffer 256 bytes
			playerID: id,
		}
		s.clients = append(s.clients, c)
		go c.write()
		s.game.State()
		if len(s.game.Players) == 2 {
			// Start game
			cardGame := CardGame()
			cardGame.shuffle()
			s.game.Event(addCardGameToStack(cardGame))
			s.game.Event(serveGame())
			s.sendState()
		}
	}
}

func (s *server) handleNewGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.game.Events = []Event{}
		players := s.game.Players
		s.game.Players = []*Player{}
		for _, player := range players {
			player.Cards = &CardStack{}
			s.game.Event(addPlayer(player))
		}
		cardGame := CardGame()
		cardGame.shuffle()
		s.game.Event(addCardGameToStack(cardGame))
		s.game.Event(serveGame())
		s.sendState()
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

func (s *server) handGetID(w http.ResponseWriter, r *http.Request) (id string, ok bool) {
	u := r.URL
	q := u.Query()
	ids, ok := q["id"]
	if !ok || len(ids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "No ID given")
		return "", false
	}
	return ids[0], true
}

func (s *server) handleData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := s.handGetID(w, r)
		if !ok {
			return
		}
		//c, ok := s.getClient(id[0])
		s.game.State()
		p, ok := s.game.Player(id)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Wrong ID")
			log.Println("wrongID:", id)
			return
		}
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		err := enc.Encode(p)
		if err != nil {
			log.Println("error handleData() encoding: ", err)
		}
	}
}

// handlePlayCard is called via a normal get request. It expects
// the id inside the url (bla.com/playcard?id=ab123)
func (s *server) handlePlayCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := s.handGetID(w, r)
		if !ok {
			return
		}
		for _, p := range s.game.Players {
			i, ok := p.Cards.find(id)
			if !ok {
				continue
			}
			s.game.Event(playCardToHeap(p, i))
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, StatusCardPlayed)
			s.sendState()
			return
		}
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, StatusCardNotFound)
	}
}

func (s *server) handleTakeCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := s.handGetID(w, r)
		if !ok {
			return
		}
		p, ok := s.game.Player(id)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, StatusPlayerNotFound)
			return
		}
		s.game.Event(takeCardFromStack(p))
		s.sendState()
	}
}

func (s *server) getClient(id string) (*client, bool) {
	for _, c := range s.clients {
		if id == c.playerID {
			return c, true
		}
	}
	return nil, false
}