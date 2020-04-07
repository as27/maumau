package main

import "github.com/google/uuid"

// Player defines the properties of a Player
type Player struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Cards *CardStack `json:"cards"`
}

func newPlayer(name string) *Player {
	return &Player{
		ID:    uuid.New().String(),
		Name:  name,
		Cards: &CardStack{},
	}
}
