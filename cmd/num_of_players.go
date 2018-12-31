package main

import (
	"fmt"
	"strconv"
)

func askNumOfPlayers() {
	fmt.Println("\nHow many opponents besides you and the dealer?")
	fmt.Printf("\n(0 - 3): ")
}

func getNumOfPlayers(s string) (int, bool) {
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
