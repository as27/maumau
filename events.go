package main

func AddPlayer(p Player) Event {
	return func(g *Game) {
		p.Cards = &CardStack{}
		g.Players = append(g.Players, p)
	}
}

func startGame() Event {
	return func(g *Game) {

		for _, p := range g.Players {
			p.Cards = &CardStack{Cards: []Card{}}
		}
		for i := 1; i <= g.NrCards; i++ {
			for _, p := range g.Players {
				p.Cards.push(g.Stack.pop())
			}
		}
	}
}

func popCardFromStack(p Player) Event {
	return func(g *Game) {
		p.Cards.push(g.Stack.pop())
	}
}

func pushCardToHeap(p Player, i int) Event {
	return func(g *Game) {
		g.Heap.push(p.Cards.take(i))
	}
}
