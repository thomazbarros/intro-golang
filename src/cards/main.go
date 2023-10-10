package main

func main() {
	//var card string = "Ace of Spades"
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println([]byte("Fim"))
}
