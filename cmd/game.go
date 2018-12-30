package main

import (
	"fmt"

	"github.com/lapubell/cards"
)

func playBlackjack() {
	fmt.Println("Welcome!")

	deck := cards.NewDeck()

	fmt.Println(deck.Cards[0].Suit.Symbol, deck.Cards[0].Type)
	// fmt.Println(deck.Cards)
}
