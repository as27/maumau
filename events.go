package main

func AddPlayer(p Player) Event {
	return func(g *Game) {
		g.Players = append(g.Players, p)
	}
}

func startGame() Event {
	return func(g *Game) {
		g.Stack = CardGame()
		g.Stack.shuffle()
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
