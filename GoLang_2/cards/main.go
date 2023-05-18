package main

import "fmt"

var a int = 5

func main() {
	cards := newDeck()
	cards.printCards()
	hand_cards, remainingCards := deal(cards, 5)
	hand_cards.printCards()
	println("-----------------------")
	remainingCards.printCards()
	println("-----------------------")
	cards.saveToFile("cards")
	cards := readFile("cards")
	cards.shuffleDeck()
	cards.printCards()

	fmt.Println(a)

}

func sxca() {
	fmt.Println(a)
}
