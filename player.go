package main

// Player defines the properties of a Player
type Player struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Cards  *CardStack `json:"cards"`
	Active bool       `json:"-"`
	Class  string     `json:"class"`
}

func newPlayer(name string) *Player {
	return &Player{
		//ID:    uuid.New().String(),
		ID:    name,
		Name:  name,
		Cards: &CardStack{},
	}
}
