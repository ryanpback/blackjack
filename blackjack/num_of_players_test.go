package blackjack

import "testing"

type playerTest struct {
	input        string
	numOfPlayers int
	errorMessage string
}

var playerTests = []playerTest{
	{"a", 0, "The number of players must be numeric"},
	{"-42", 0, "You must input a positive number of players"},
	{"0", 0, "This game requires at least one dealer and one player"},
	{"1", 0, "This game requires at least one dealer and one player"},
	{"2", 2, ""},
	{"3", 3, ""},
	{"4", 4, ""},
	{"5", 5, ""},
	{"6", 6, "Too many players"},
}

func TestUserCantGiveBadDataForNumOfPlayers(t *testing.T) {
	for _, pt := range playerTests {
		err := SetNumberOfPlayers(pt.input)

		if err != nil {
			if err.Error() != pt.errorMessage {
				t.Errorf("Wrong error message sent. Expected %s, got %s", pt.errorMessage, err.Error())
			}
		}
	}
}
