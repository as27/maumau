package main

// Game defines the logic of the maumau game
type Game struct {
	GameState
	Events     []Event `json:"-"`
	RedoEvents []Event `json:"-"`
	NrCards    int     `json:"nr_cards,omitempty"`
}

// GameState includes all properties of the game, which are having
// a state. The is just changed by an event.
type GameState struct {
	Stack        *CardStack `json:"stack,omitempty"`
	Heap         *CardStack `json:"heap,omitempty"`
	HeapHead     Card       `json:"heap_head,omitempty"`
	Players      []*Player  `json:"players,omitempty"`
	ActivePlayer int        `json:"active_player"`
}

func newGame() *Game {
	return &Game{
		GameState: GameState{
			Stack:        &CardStack{},
			Heap:         &CardStack{},
			ActivePlayer: 0,
		},
		NrCards: 6,
	}
}

// Event takes an event and adds that to the game
// to get a new state Init() has to be called
func (g *Game) Event(e Event) {
	// every new event clears the redo slice
	g.RedoEvents = []Event{}
	g.Events = append(g.Events, e)
}

// Init clears the internal state
func (g *Game) Init() {
	g.Stack = &CardStack{}
	g.Heap = &CardStack{}
	g.Players = []*Player{}
}

// State executes all events and creates a new state for
// the Stack, Heap and the Players
func (g *Game) State() {
	g.Init()
	for _, e := range g.Events {
		e(g)
	}
	g.HeapHead = g.Heap.peek()
}

// Player returns the player matching to the given ID.
// when no ID is availiable a nil pointer is returned.
func (g *Game) Player(id string) (*Player, bool) {
	for _, p := range g.Players {
		if id == p.ID {
			return p, true
		}
	}
	return nil, false
}

// NextPlayer takes the current player ID and returns the
// next player at the table. If there is no ID a nil pointer
// is returned
func (g *Game) NextPlayer(id string) (*Player, bool) {
	found := -1
	// index of the current player
	for i, p := range g.Players {
		if id == p.ID {
			found = i
			break
		}
	}
	if found == -1 {
		return nil, false
	}
	next := found + 1
	if len(g.Players)-1 == found {
		// if current player is the last in the slice
		// the first player is next
		next = 0
	}
	return g.Players[next], true
}

// SetActivePlayer takes an ID to set the active player
// all other players will set es not active.
// if there is no ID, false ist returned
func (g *Game) SetActivePlayer(id string) bool {
	// first check if there is a player to the id
	// if not nothing should be changed.
	_, ok := g.Player(id)
	if !ok {
		return false
	}
	for _, p := range g.Players {
		if id == p.ID {
			p.Active = true
		} else {
			p.Active = false
		}
	}
	return true
}

// Event is a function, which takes a pointer to the game
// When the game calculates the state all events are called
type Event func(g *Game)
