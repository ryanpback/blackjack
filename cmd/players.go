package main

import (
	"fmt"

	"github.com/lapubell/cards"
)

type player struct {
	name string
	hand []cards.Card
	// wallet int - save for betting later?
}

// CreatePlayers will always create at least the Dealer and the User's Player. It will create more depending on the number chosen by the user
func createPlayers(num int) []player {
	computerPlayerNames := [...]string{"Crucio", "Imperius", "Avada Kedavra"}

	players := []player{
		{"Dealer", []cards.Card{}},
	}

	usersPlayer := createUsersPlayer()
	players = append(players, usersPlayer)

	for i := 1; i <= num; i++ {
		players = append(players, player{computerPlayerNames[i-1], []cards.Card{}})
	}

	return players
}

func createUsersPlayer() player {
	fmt.Printf("Enter the name you'd like for your player: ")
	userInput := getUserInput()

	player := player{userInput, []cards.Card{}}

	return player
}
