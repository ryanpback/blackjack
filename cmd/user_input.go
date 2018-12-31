package main

import "fmt"

func getUserInput() string {
	var input string
	fmt.Scanln(&input)

	return input
}
