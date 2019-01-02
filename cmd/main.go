package main

import "fmt"

func main() {
	readyToPlay := areWePlaying()

	if readyToPlay {
		players := setUpGame()
		playBlackjack(players)
	}

	fmt.Println("Goodbye")
}
