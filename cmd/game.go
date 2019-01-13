package main

import (
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"

	"github.com/ryanpback/blackjack/blackjack"
	"github.com/ryanpback/blackjack/player"
	"github.com/ryanpback/blackjack/userInput"
)

func setUpPlayers() {
	data := struct {
		Min int
		Max int
	}{
		Min: blackjack.MinNumberOfPlayers,
		Max: blackjack.MaxNumberOfPlayers,
	}
	t := template.Must(template.New("askNumPlayers").Parse(userInput.AskNumberOfPlayers))
	t.Execute(os.Stdout, data)

	for {
		if blackjack.NumberOfPlayers > 0 {
			break
		}
		userInput := userInput.GetUserInput()
		err := blackjack.SetNumberOfPlayers(userInput)

		if err != nil {
			callClear()
			fmt.Println(err.Error())
			t.Execute(os.Stdout, data)
		}
	}

}

func playBlackjack(players []player.Player) {
	for {
		callClear()
		fmt.Println(userInput.GameStart)

		drawTable(players)
		blackjack.SetPlayers(players)
		state := blackjack.CheckGameState()

		if !state {
			break
		}
	}
}

func drawTable(players []player.Player) {
	// lets populate the default table object with some empty strings, then we
	// can draw each player in place
	data := make(map[string]string, 18)
	data["P01"] = "     "
	data["P02"] = "     "
	data["P03"] = "     "
	data["P11"] = "     "
	data["P12"] = "     "
	data["P13"] = "     "
	data["P21"] = "     "
	data["P22"] = "     "
	data["P23"] = "     "
	data["P31"] = "     "
	data["P32"] = "     "
	data["P33"] = "     "
	data["P41"] = "     "
	data["P42"] = "     "
	data["P43"] = "     "
	data["P51"] = "     "
	data["P52"] = "     "
	data["P53"] = "     "

	// populate the layout
	for i := range players {
		if i > len(players) {
			break
		}
		currentPlayerIndex := strconv.Itoa(i)
		playerAvatar := strings.Split(players[i].Draw(), "\n")
		data["P"+currentPlayerIndex+"1"] = playerAvatar[0]
		data["P"+currentPlayerIndex+"2"] = playerAvatar[1]
		data["P"+currentPlayerIndex+"3"] = playerAvatar[2]
	}

	t := template.Must(template.New("drawTable").Parse(layout))
	t.Execute(os.Stdout, data)
}
