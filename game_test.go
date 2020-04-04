package main

import (
	"reflect"
	"testing"
)

func TestGameEvents(t *testing.T) {
	g := newGame()
	g.Shuffle()
	g.Event(startGame())
	g.Event(AddPlayer(Player{Name: "Max"}))
	g.Event(AddPlayer(Player{Name: "Maja"}))
	g.State()
	if g.Players[0].Name != "Max" {
		t.Errorf("first player should be Max: got %#v", g.Players[0])
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
