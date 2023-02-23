package main

import "fmt"

// create type of deck which is a slice of stings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues:= []string{"Ace", "Two", "Three", "Four"}

	for _, s := range cardSuits {
		for _, v := range cardValues {
			card := v + " of " + s
			cards = append(cards, card)
		}
	}

	return cards
}

// receiver for any type deck
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}