package main

import "testing"

func TestCreateNumberOfUsersPassedToFunc(t *testing.T) {
	numOfPlayersOptions := [...]int{0, 1, 2, 3}

	for i := 0; i < len(numOfPlayersOptions)-1; i++ {
		players := createPlayers(numOfPlayersOptions[i])

		if len(players) != numOfPlayersOptions[i]+2 {
			t.Errorf("Passing in %v to createPlayers should result in x + 2 players", numOfPlayersOptions[i])
		}
	}
}
