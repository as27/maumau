package main

func addCardGameToStack(cs *CardStack) Event {
	return func(g *Game) {
		g.Stack.Cards = append(cs.Cards, g.Stack.Cards...)
	}
}

func addPlayer(p *Player) Event {
	return func(g *Game) {
		p.Cards = &CardStack{}
		g.Players = append(g.Players, p)
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
	}
}

func popCardFromStack(p *Player) Event {
	return func(g *Game) {
		p.Cards.push(g.Stack.pop())
	}
}

func pushCardToHeap(p *Player, i int) Event {
	return func(g *Game) {
		g.Heap.push(p.Cards.take(i))
	}
}
