package main

import (
	"fmt"

	"github.com/lapubell/cards"
)

func playBlackjack() {
	fmt.Println("Welcome!")

	deck := cards.GetDeck()

	fmt.Println(deck)
}
