package main

import (
	"fmt"
	"strings"
	"time"
	"os"
	"math/rand"
)

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

// d arg with type deck, handSize arg type int
func deal(d deck, handSize int) (deck, deck) {
	return d [:handSize], d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// option 1 - log error and return a call to newDeck
		// option 2 - log error and quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	/* this was a little confusing */
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}