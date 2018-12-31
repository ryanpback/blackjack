package main

import (
	"fmt"
	"strconv"
)

func getNumberOfPlayers() int {
	// Keep numOfPlayers as unreachable value
	numOfPlayers := 99
	correctInput := false

	for {
		askNumOfPlayers()
		userInput := getUserInput()
		numOfPlayers, correctInput = getPlayerCount(userInput)

		if correctInput {
			return numOfPlayers
		}

		if !correctInput && numOfPlayers != 99 {
			fmt.Printf("\n'%v' is not in the correct range", userInput)
		}

		if !correctInput && numOfPlayers == 99 {
			fmt.Printf("\n'%v' is not a valid option", userInput)
		}
	}
}

func askNumOfPlayers() {
	fmt.Println("\nHow many players besides you and the dealer?")
	fmt.Printf("\n(0 - 3): ")
}

func getPlayerCount(s string) (int, bool) {
	i, err := strconv.Atoi(s)

	// Can't convert string value to int
	if err != nil {
		return 99, false
	}

	if i >= 0 && i <= 3 {
		return i, true
	}

	return i, false
}
