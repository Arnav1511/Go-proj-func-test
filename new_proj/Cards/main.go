package main

import (
	"fmt"
)

var cards = "Ace of Spades"

func main() {
	card := "Ace of Spades"
	card = "King of Hearts" // Example of modifying the card variable
	fmt.Println(card)
	fmt.Println(cards) // Print the global cards variable
}
