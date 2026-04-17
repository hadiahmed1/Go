package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("deck")
	cards := newDeckFromFile("deck")
	cards.print()
	cards.shuffle()
	cards.print()
}
