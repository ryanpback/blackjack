package userInput

import "fmt"

func GetUserInput() string {
	var input string
	fmt.Scanln(&input)

	return input
}
