package main

import (
	"fmt"
	"strings"

	"github.com/ryanpback/blackjack/userInput"
)

func displayWelcome() {
	fmt.Println(userInput.Intro)
}

func displayGoodbye() {
	fmt.Println(userInput.Outro)
}

func closeGame() {
	// maybe we clear the screen here?
	// callClear()
	displayGoodbye()
}

func playOrQuit(s string) (bool, bool) {
	letter := strings.ToLower(s)

	// if they just hit enter, let's give them the default value
	if letter == "" {
		letter = "y"
	}

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
		readyToPlay, correctInput = playOrQuit(userInput.GetUserInput())

		if readyToPlay && correctInput {
			callClear()
			return true
		}

		if !readyToPlay && correctInput {
			return false
		}

		fmt.Println(userInput.InvalidOption)
	}
}
