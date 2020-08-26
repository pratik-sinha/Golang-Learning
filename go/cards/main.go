package main

func main() {
	//var card string = "Ace Of Spades"
	//card := "Ace Of Spades"
	//card := newCard()
	//cards = append(cards, "Six of Spades")
	//append generates a new slice
	// cards := newDeck()
	// hand, remainingDeck := cards.deal(3)
	// hand.print()
	// remainingDeck.print()
	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
}
