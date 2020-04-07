package main

import "github.com/google/uuid"

func addCardGameToStack(cs *CardStack) Event {
	return func(g *Game) {
		g.Stack.Cards = append(cs.Cards, g.Stack.Cards...)
	}
}

func addPlayer(p *Player) Event {
	return func(g *Game) {
		p.Cards = &CardStack{}
		if p.ID == "" {
			p.ID = uuid.New().String()
		}
		g.Players = append(g.Players, p)
	}
}

func nextPlayer() Event {
	return func(g *Game) {
		g.ActivePlayer++
		if g.ActivePlayer > len(g.Players)-1 {
			g.ActivePlayer = 0
		}
	}
}

// serveGame serves the cards from the stack to the players
func serveGame() Event {
	return func(g *Game) {
		// a new emtpy hand for every player
		for _, p := range g.Players {
			p.Cards = &CardStack{Cards: []Card{}}
		}
		for i := 1; i <= g.NrCards; i++ {
			for j := range g.Players {
				g.Players[j].Cards.push(g.Stack.pop())
			}
		}
		g.Heap.push(g.Stack.pop())
	}
}

func takeCardFromStack(p *Player) Event {
	return func(g *Game) {
		p.Cards.push(g.Stack.pop())
	}
}

func playCardToHeap(p *Player, i int) Event {
	return func(g *Game) {
		g.Heap.push(p.Cards.take(i))
	}
}
