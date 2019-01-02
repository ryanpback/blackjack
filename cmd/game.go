package main

import (
	"fmt"

	"github.com/lapubell/cards"
)

func playBlackjack(players []player) {
	fmt.Println("Let's play!")
	fmt.Println("========================")

	deck := cards.NewDeck()

	fmt.Println(deck.Cards[0].Suit.Symbol, deck.Cards[0].Type)
}
