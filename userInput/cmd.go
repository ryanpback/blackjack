// Package userInput is an abstraction package so we have a common spot to put
// all functions that need to take in user input. More stuff will be moved here
// as the main package gets slimmer, keeping all speration of concerns happy.
package userInput

import "fmt"

// GetPlayerInformationCMD this function will get input from the user to build
// either an automated (computer) player or a human controlled one.
func GetPlayerInformationCMD() (string, bool, int) {
	var inputName string
	var inputHuman bool
	var inputWallet int

	var isHuman string
	fmt.Println(PlayerIsHuman)
	fmt.Scanln(&isHuman)
	inputHuman = validateInputForHuman(isHuman)
	if !inputHuman {
		// this is a computer controlled user, so let's set some defaults
		return "Computer Player", false, 10000
	}

	fmt.Println(getPlayerName)
	fmt.Scanln(&inputName)

	return inputName, inputHuman, inputWallet
}
