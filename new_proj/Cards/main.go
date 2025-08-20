package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamond", newCard()}
	cards = append(cards, "Six of Spades")

	// Print the cards in hand
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
