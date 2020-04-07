package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// CardGame defines all availiable cards inside the game.
// This function is called to generate a stack for the game.
// Exept for the wish cards, every card needs a value and a color.
func CardGame() *CardStack {
	var cs CardStack
	colors := []string{"c1", "c2", "c3", "c4"}
	values := []string{"1", "1", "2", "2", "3", "3", "4", "4", "5", "5"}
	for _, c := range colors {
		for _, v := range values {
			newCard := Card{
				ID:          uuid.New().String(),
				Color:       c,
				Value:       v,
				SkipPlayers: 0,
				TakeN:       0,
				WishColor:   false,
			}
			cs.push(newCard)
		}
		// every color has a skip next player
		newCard := Card{
			ID:          uuid.New().String(),
			Color:       c,
			Value:       "->",
			SkipPlayers: 1,
		}
		cs.push(newCard)
		// every color has a 2+ card
		newCard = Card{
			ID:    uuid.New().String(),
			Color: c,
			Value: "+2",
			TakeN: 2,
		}
		cs.push(newCard)
		// wish cards don't need a color or value
		newCard = Card{
			ID:        uuid.New().String(),
			Color:     "wish",
			WishColor: true,
		}
		cs.push(newCard)
	}
	return &cs
}

// Card defines the propperties of a card.
type Card struct {
	ID          string `json:"id"`
	Color       string `json:"color"`
	Value       string `json:"value"` // number or name e.g. 1, K, J
	SkipPlayers int    `json:"skip_players"`
	TakeN       int    `json:"take_n"`     // next player has to take this nr of cards
	WishColor   bool   `json:"wish_color"` // defines a wish card
}

// Check validates the next cards due to the rules of the game.
// In maumau the next card needs the same color or value. Wish cards
// are allowed to play on every card.
func (c Card) Check(next Card) bool {
	switch {
	case next.WishColor:
		return true
	case c.Color == next.Color:
		return true
	case c.Value == next.Value:
		return true
	}
	return false
}

// CardStack defines a group of cards. The heap, the stack and the cards
// a player has on his hand are defines as CardStack.
type CardStack struct {
	Cards []Card `json:"cards"`
}

func (cs *CardStack) push(c Card) {
	cs.Cards = append(cs.Cards, c)
}

func (cs *CardStack) pop() Card {
	var c Card
	if len(cs.Cards) == 0 {
		return c
	}
	c = cs.Cards[len(cs.Cards)-1]
	cs.Cards = cs.Cards[:len(cs.Cards)-1]
	return c
}

func (cs *CardStack) popAllowed(c Card) bool {
	return cs.peek().Check(c)
}

func (cs *CardStack) peek() Card {
	var c Card
	if len(cs.Cards) == 0 {
		return c
	}
	return cs.Cards[len(cs.Cards)-1]
}

// find returns the index of the card with a given
// id. When that card is not inside the stack -1
// is returned. To check the success use the second
// ok parameter
func (cs *CardStack) find(id string) (int, bool) {
	for i, c := range cs.Cards {
		if c.ID == id {
			return i, true
		}
	}
	return -1, false
}

// take returns the i element of the CardStack and removes
// that from the stack.
func (cs *CardStack) take(i int) Card {
	c := cs.Cards[i]
	cs.Cards = append(cs.Cards[:i], cs.Cards[i+1:]...)
	return c
}

func (cs *CardStack) len() int {
	return len(cs.Cards)
}

func (cs *CardStack) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cs.Cards), func(i, j int) {
		cs.Cards[i], cs.Cards[j] = cs.Cards[j], cs.Cards[i]
	})
}
