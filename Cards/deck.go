package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	d := deck{}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cards := []string{"Ace", "2", "2", "4", "5", "6", "7", "8", "9", "10", "K", "Q", "J"}
	for _, suit := range suits {
		for _, card := range cards {
			d = append(d, (card + " of " + suit))
		}
	}
	return d
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// d is the reciver
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	str := d.toString()
	return os.WriteFile(filename, []byte(str), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	str := string(bs)
	var newDeck deck = strings.Split(str, ",")
	return newDeck
}

func (d deck) shuffle() {
	for i := range d {
		ri := rand.Intn(len(d))
		// swap i & ri
		d[i], d[ri] = d[ri], d[i]
	}
}
