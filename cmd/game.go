package main

import (
	"fmt"

	"github.com/lapubell/cards"
)

func playBlackjack() {
	fmt.Println("Let's play!")
	fmt.Println("========================")

	numOfPlayers := getNumberOfPlayers()
	fmt.Println("\nnum: ", numOfPlayers)

	deck := cards.NewDeck()
	fmt.Println(deck.Cards[0].Suit.Symbol, deck.Cards[0].Type)
}
