package main

type Game struct {
	Stack         *CardStack
	Heap          *CardStack
	Players       []*Player
	InitialEvents []Event
	Events        []Event
	NrCards       int
}

func newGame() *Game {
	return &Game{
		Stack:   &CardStack{},
		Heap:    &CardStack{},
		NrCards: 6,
	}
}

func (g *Game) Event(e Event) {
	g.Events = append(g.Events, e)
}
func (g *Game) InitialEvent(e Event) {
	g.InitialEvents = append(g.InitialEvents, e)
}

func (g *Game) Init() {
	g.Stack = &CardStack{}
	g.Heap = &CardStack{}
	g.Players = []*Player{}

	for _, e := range g.InitialEvents {
		e(g)
	}
}

func (g *Game) State() {
	for _, e := range g.Events {
		e(g)
	}
}

type Event func(g *Game)

type Player struct {
	ID    string
	Name  string
	Cards *CardStack
}
