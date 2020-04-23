package main

import (
	"reflect"
	"testing"
)

func TestCard_Check(t *testing.T) {
	type args struct {
		next Card
	}
	tests := []struct {
		name string
		c    Card
		args args
		want bool
	}{
		{
			"same color",
			Card{Color: "red", Value: "10"},
			args{next: Card{Color: "red", Value: "1"}},
			true,
		},
		{
			"same value",
			Card{Color: "red", Value: "10"},
			args{next: Card{Color: "green", Value: "10"}},
			true,
		},
		{
			"not allowed",
			Card{Color: "red", Value: "10"},
			args{next: Card{Color: "green", Value: "5"}},
			false,
		},
		{
			"wish card",
			Card{Color: "red", Value: "10"},
			args{next: Card{Color: "", WishColor: true}},
			true,
		},
		{
			"nach wish card",
			Card{Color: "", Value: "", WishColor: true},
			args{next: Card{Color: "green", Value: "5", WishColor: false}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Check(tt.args.next); got != tt.want {
				t.Errorf("Card.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardStackMethods(t *testing.T) {
	cards := []Card{
		Card{ID: "1"},
		Card{ID: "2"},
		Card{ID: "3"},
		Card{ID: "4"},
	}
	// Add all the cards to the stack
	cs := CardStack{}
	for _, c := range cards {
		cs.push(c)
	}
	// check the cards inside the stack
	if !reflect.DeepEqual(cards, cs.Cards) {
		t.Errorf("Slices should be equal:\ngot:\n%v\nwant:\n%v", cs.Cards, cards)
	}
	// take the last card of the stack
	got := cs.pop()
	if got != cards[3] {
		t.Errorf("pop(): got:%v\nwant:%v", got, cards[3])
	}
	// check the size of the stack
	if cs.len() != 3 {
		t.Errorf("Size of the stack is %d\nwant: 3", cs.len())
	}
	// check the last card with peek
	got = cs.peek()
	if got != cards[2] {
		t.Errorf("peek(): got:%v\nwant:%v", got, cards[3])
	}
}
