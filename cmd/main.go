package main

import (
	"fmt"
)

func main() {
	readyToPlay := false
	correctInput := false

	for {
		displayWelcome()
		userInput := getUserInput()
		readyToPlay, correctInput = playOrQuit(userInput)

		if readyToPlay && correctInput {
			playBlackjack()
			break
		}

		if !readyToPlay && correctInput {
			fmt.Println("Goodbye")
			break
		}

		fmt.Println("\n'" + userInput + "' is not a valid option")
	}
}
