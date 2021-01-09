package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards.print()
	cards.push("50 shades of grey")
	cards.print()
}

func newCard() string {
	return "Five of Diamond"
}
