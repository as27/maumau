package main

// Game defines the logic of the maumau game
type Game struct {
	Stack    *CardStack `json:"stack,omitempty"`
	Heap     *CardStack `json:"heap,omitempty"`
	HeapHead Card       `json:"heap_head,omitempty"`
	Players  []*Player  `json:"players,omitempty"`
	Events   []Event    `json:"-"`
	NrCards  int        `json:"nr_cards,omitempty"`
}

func newGame() *Game {
	return &Game{
		Stack:   &CardStack{},
		Heap:    &CardStack{},
		NrCards: 6,
	}
}

// Event takes an event and adds that to the game
// to get a new state Init() has to be called
func (g *Game) Event(e Event) {
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

// Event is a function, which takes a pointer to the game
// When the game calculates the state all events are called
type Event func(g *Game)
