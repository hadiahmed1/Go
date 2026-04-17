package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// testing length
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	// testing card values
	if d[0] != "Ace of Hearts" {
		t.Error("1st card was expected to be Ace of Hearts, got", d[0])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	df := newDeckFromFile("_decktesting")

	for i, card := range d {
		if df[i] != card {
			t.Errorf("Deck not saved to file correctly, card mistmatch")
		}
	}
	os.Remove("_decktesting")
}
