package main

// Game defines the logic of the maumau game
type Game struct {
	Stack    *CardStack
	Heap     *CardStack
	HeapHead Card
	Players  []*Player
	Events   []Event
	NrCards  int
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

// Event is a function, which takes a pointer to the game
// When the game calculates the state all events are called
type Event func(g *Game)
