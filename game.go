package main

type Game struct {
	Stack   *CardStack
	Heap    *CardStack
	Players []Player
	Events  []Event
	NrCards int
}

func newGame() *Game {
	return &Game{
		Stack: &CardStack{},
		Heap:  &CardStack{},
	}
}

func (g *Game) Event(e Event) {
	g.Events = append(g.Events, e)
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
