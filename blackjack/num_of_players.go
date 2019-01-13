package blackjack

import (
	"errors"
	"strconv"
)

// MinNumberOfPlayers this is the min number of players for this game
// (dealer and player)
var MinNumberOfPlayers = 2

// MaxNumberOfPlayers this is the max number of players for this game
// up to 5 additional players (5 regular and 1 dealer)
var MaxNumberOfPlayers = 6

// NumberOfPlayers this is the number of players in the current game
var NumberOfPlayers int

// SetNumberOfPlayers get the number of players from the user. Return
func SetNumberOfPlayers(playerInput string) error {
	numOfPlayers, err := strconv.Atoi(playerInput)

	// Can't convert string value to int
	if err != nil {
		return errors.New("The number of players must be numeric")
	}

	if numOfPlayers < 0 {
		return errors.New("You must input a positive number of players")
	}

	if numOfPlayers < MinNumberOfPlayers {
		return errors.New("This game requires at least one dealer and one player")
	}

	if numOfPlayers > MaxNumberOfPlayers {
		return errors.New("Too many players")
	}

	NumberOfPlayers = numOfPlayers
	return nil
}
