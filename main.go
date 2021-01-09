package main

func main() {
	cards := newDeck()
	hand, remainningCards := deal(cards, 3)
	hand.print()
	remainningCards.print()
}
