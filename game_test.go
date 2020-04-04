package main

import "testing"

func TestGameEvents(t *testing.T) {
	g := newGame()
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

}
