package main

import "github.com/ryanpback/blackjack/blackjack"

func main() {
	readyToPlay := areWePlaying()

	if readyToPlay {
		setUpPlayers()
		players := blackjack.CreatePlayers("cmd")
		playBlackjack(players)
	}

	closeGame()
}
