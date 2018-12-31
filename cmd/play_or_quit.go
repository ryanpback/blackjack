package main

import (
	"fmt"
	"strings"
)

func displayWelcome() {
	fmt.Println("\nReady to play BlackJack?")
	fmt.Println("========================")
	fmt.Printf("\nYes (Y) | Quit (Q): ")
}

func playOrQuit(s string) (bool, bool) {
	letter := strings.ToLower(s)

	if letter == "y" {
		return true, true
	}
	if letter == "q" {
		return false, true
	}
	return false, false
}

func areWePlaying() bool {
	readyToPlay := false
	correctInput := false

	for {
		displayWelcome()
		userInput := getUserInput()
		readyToPlay, correctInput = playOrQuit(userInput)

		if readyToPlay && correctInput {
			return true
		}

		if !readyToPlay && correctInput {
			return false
		}

		fmt.Println("\n'" + userInput + "' is not a valid option")
	}
}
