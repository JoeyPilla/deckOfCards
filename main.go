package main

import (
	"fmt"

	"./deck"
)

func main() {
	d := deck.NewDeck()
	fmt.Println(d.Deck)
}
