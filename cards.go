package main

import (
	"math/rand"
	"time"
)

func CardGame() *CardStack {
	var cs CardStack
	colors := []string{"rot", "gelb", "grün", "blau"}
	values := []string{"1", "1", "2", "2", "3", "3", "4", "4", "5", "5"}
	for _, c := range colors {
		for _, v := range values {
			newCard := Card{
				Color:       c,
				Value:       v,
				SkipPlayers: 0,
				TakeN:       0,
				WishColor:   false,
			}
			cs.push(newCard)
		}
		newCard := Card{
			Color:       c,
			SkipPlayers: 1,
		}
		cs.push(newCard)
		newCard = Card{
			Color: c,
			TakeN: 2,
		}
		cs.push(newCard)
		newCard = Card{
			WishColor: true,
		}
		cs.push(newCard)
	}
	return &cs
}

type Card struct {
	Color       string // Farbe der Karte
	Value       string // Zahl der Karte
	SkipPlayers int    // Anzahl der Spieler die Überspungen werden
	TakeN       int    // Anzahl der Karten die der nächste Spieler nehmen muss
	WishColor   bool
}

type CardStack struct {
	Cards []Card
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

func (cs *CardStack) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cs.Cards), func(i, j int) {
		cs.Cards[i], cs.Cards[j] = cs.Cards[j], cs.Cards[i]
	})
}
