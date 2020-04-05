package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGameEvents(t *testing.T) {

	g := newGame()
	cardGame := CardGame()
	g.Event(addCardGameToStack(cardGame))
	g.Event(addPlayer(newPlayer("Max")))
	g.Event(addPlayer(newPlayer("Maja")))
	g.Event(serveGame())
	g.State() // State can be called more then once
	g.State()
	if g.Stack.len() == 0 {
		t.Error("stack is empty some cards should be added to the game")
	}
	if g.Players[0].Name != "Max" {
		t.Errorf("first player should be Max: got %#v", g.Players[0])
	}
	if g.Players[0].ID == "" {
		fmt.Println(g.Players[0].ID)
		t.Error("Player 0 has no ID")
	}
	if g.Players[0].Cards.len() != g.NrCards {
		t.Errorf("Player 0 had %d cards! Expect: %d", g.Players[0].Cards.len(), g.NrCards)
		t.Errorf("%#v", g.Players[0].Cards)
	}

	// Player 0 has to take the top card of the stack
	topCard := g.Stack.peek()
	g.Event(takeCardFromStack(g.Players[0]))
	g.State()
	pTopCard := g.Players[0].Cards.peek()
	if !reflect.DeepEqual(topCard, pTopCard) {
		t.Error("popCardFromStack does not serve the right card")
	}
	// Player 0 plays one card to the heap
	pFirstCard := g.Players[0].Cards.Cards[0]
	g.Event(playCardToHeap(g.Players[0], 0))
	g.State()
	hTopCard := g.Heap.pop()
	g.Heap.push(hTopCard)
	if pFirstCard.ID == "" {
		t.Error("pFirstCard has no ID")
	}
	if pFirstCard.ID != hTopCard.ID {
		t.Error("pushCardToHeap does not push the right card to the heap")
	}
}
