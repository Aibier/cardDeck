package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func cardName(name string) string {
	return name
}