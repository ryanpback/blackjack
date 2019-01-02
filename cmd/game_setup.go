package main

func setUpGame() []player {
	numOfPlayers := getNumberOfPlayers()

	return createPlayers(numOfPlayers)
}
