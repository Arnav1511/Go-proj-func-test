package main

func main() {

	cards := deck{"Ace of Diamond", newCard()}
	cards = append(cards, "Six of Spades")

	// Print the cards in hand
	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
