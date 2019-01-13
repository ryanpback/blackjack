package blackjack

import (
	"fmt"

	"github.com/lapubell/cards"
)

func setupAndDeal() {
	fmt.Println("Let's start. Shuffle time!")
	deck = cards.NewDeck()
	// shuffle 3 times, that's how we roll baybee!
	deck.Shuffle()
	deck.Shuffle()
	deck.Shuffle()
	fmt.Println("Now we deal the cards")
	fmt.Println("Press [Enter] to continue:")
	fmt.Scanln()
	gameState = "cardsDelt"
}
