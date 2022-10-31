package main

func main() {

	/* cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") */

	/* cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print() */

	/* greeting := "Hi there!"
	fmt.Println([]byte(greeting)) */

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// fmt.Println(cards.toString())

}

/* func newCard() string {
	return "Five of Diamonds"
} */
