package main

import "github.com/google/uuid"

// Player defines the properties of a Player
type Player struct {
	ID    string
	Name  string
	Cards *CardStack
}

func newPlayer(name string) *Player {
	return &Player{
		ID:    uuid.New().String(),
		Name:  name,
		Cards: &CardStack{},
	}
}
