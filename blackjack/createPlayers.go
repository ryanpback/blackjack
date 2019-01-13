package blackjack

import (
	"github.com/ryanpback/blackjack/player"
	"github.com/ryanpback/blackjack/userInput"
)

// CreatePlayers will run through the steps necessary to create all the game
// players. The inputType argument is so that this package knows what type of
// interface to use to work with the user. (cmd|http|?)
func CreatePlayers(inputType string) []player.Player {
	players := make([]player.Player, NumberOfPlayers)

	if inputType == "cmd" {
		userInput.ShowCreateDealerMessageCMD()
		players[0] = player.SetupPlayer("Dealer", false, 0)

		for i := range players {
			if i == 0 { // skip the dealer, that will always be the first player
				continue
			}

			players[i] = player.SetupPlayer(userInput.GetPlayerInformationCMD())
		}
	} else {
		// not sure how we want to handle the other interfaces...
		panic("this isn't handled yet...")
	}

	return players
}
