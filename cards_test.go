package main

import "testing"

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
			args{next: Card{Color: "green", WishColor: true}},
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
