package blackjack

import (
	"github.com/lapubell/cards"
	"github.com/ryanpback/blackjack/player"
)

var deck cards.Deck
var players []player.Player
var gameState string

// SetPlayers populate the game players
func SetPlayers(p []player.Player) {
	players = p
}
