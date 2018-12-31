package main

import (
	"fmt"

	"github.com/lapubell/cards"
)

func playBlackjack() {
	fmt.Println("Let's play!")
	fmt.Println("========================")

	// Keep numOfPlayers as unreachable value
	numOfPlayers := 99
	correctInput := false
	var input string
	for {
		askNumOfPlayers()
		fmt.Scanln(&input)
		numOfPlayers, correctInput = getNumOfPlayers(input)

		if correctInput {
			fmt.Println("Why am I here")
			break
		}

		if !correctInput && numOfPlayers != 99 {
			fmt.Printf("\n'%v' is not in the correct range", input)
		}

		if !correctInput && numOfPlayers == 99 {
			fmt.Printf("\n'%v' is not a valid option", input)
		}
	}

	fmt.Println("\nnum: ", numOfPlayers)
	deck := cards.NewDeck()

	fmt.Println(deck.Cards[0].Suit.Symbol, deck.Cards[0].Type)
	// fmt.Println(deck.Cards)
}
