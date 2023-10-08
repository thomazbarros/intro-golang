package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println([]byte("Fim"))
}
