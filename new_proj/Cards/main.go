package main

func main() {

	cards := newDeck()

	// Print the cards in hand
	cards.print()

	//Dealing cards
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
