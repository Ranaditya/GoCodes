package main

import (
	"fmt"
)

type card struct {
	value string
	suit  string
}

func main() {

	/*cards := newDeck()
	//save to file
	cards.saveToFile("my_deck")
	//read from file
	newDeck := newDeckFromFile("my_deck")

	newDeck.shuffle()
	newDeck.print()
	*/
	c1 := card{suit: "Spades", value: "Ace"}
	fmt.Println(c1)
}
