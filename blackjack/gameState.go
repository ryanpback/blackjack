package blackjack

// CheckGameState determine what state the game is in, and call the necessary
// state functions to progress in game play. as long as we return true from this
// function, gameplay will continue
func CheckGameState() bool {
	if gameState == "" {
		setupAndDeal()
	}

	return false
}
