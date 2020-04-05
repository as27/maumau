package main

import (
	"reflect"
	"testing"
)

func TestGameEvents(t *testing.T) {
	makePlayer := func(name string) *Player {
		return &Player{Name: name, Cards: &CardStack{}}
	}
	g := newGame()
	cardGame := CardGame()
	g.InitialEvent(addCardGameToStack(cardGame))
	g.InitialEvent(addPlayer(makePlayer("Max")))
	g.InitialEvent(addPlayer(makePlayer("Maja")))
	g.InitialEvent(serveGame())
	g.Init()
	g.State()
	g.State()
	if g.Players[0].Name != "Max" {
		t.Errorf("first player should be Max: got %#v", g.Players[0])
	}
	if g.Players[0].Cards.len() != g.NrCards {
		t.Errorf("Player 0 had %d cards! Expect: %d", g.Players[0].Cards.len(), g.NrCards)
		t.Errorf("%#v", g.Players[0].Cards)
	}
	if g.Stack.len() == 0 {
		t.Error("stack is empty some cards should be added to the game")
	}
	topCard := g.Stack.pop()
	g.Stack.push(topCard)
	g.Event(popCardFromStack(g.Players[0]))
	g.State()
	pTopCard := g.Players[0].Cards.pop()
	g.Players[0].Cards.push(pTopCard)
	if !reflect.DeepEqual(topCard, pTopCard) {
		t.Error("popCardFromStack does not serve the right card")
	}
	pFirstCard := g.Players[0].Cards.Cards[0]
	g.Event(pushCardToHeap(g.Players[0], 0))
	g.State()
	hTopCard := g.Heap.pop()
	g.Heap.push(hTopCard)
	if !reflect.DeepEqual(pFirstCard, hTopCard) {
		t.Error("pushCardToHeap does not push the right card to the heap")
	}
}
