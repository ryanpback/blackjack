package userInput

import "strings"

func validateInputForHuman(input string) bool {
	input = strings.ToLower(input)
	if input == "" || input == "y" {
		return true
	}

	return false
}
