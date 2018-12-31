package main

import (
	"fmt"
)

func main() {
	readyToPlay := areWePlaying()

	if readyToPlay {
		playBlackjack()
	}

	fmt.Println("Goodbye")
}
