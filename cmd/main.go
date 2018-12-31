package main

import (
	"fmt"
)

func main() {
	var input string
	readyToPlay := false
	correctInput := false

	for {
		displayWelcome()
		fmt.Scanln(&input)
		readyToPlay, correctInput = playOrQuit(input)

		if readyToPlay && correctInput {
			playBlackjack()
			break
		}

		if !readyToPlay && correctInput {
			fmt.Println("Goodbye")
			break
		}

		fmt.Println("\n'" + input + "' is not a valid option")
	}
}
