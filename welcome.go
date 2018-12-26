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
