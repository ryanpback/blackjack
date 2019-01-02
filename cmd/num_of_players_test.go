package main

import "testing"

type playerTest struct {
	input        string
	numOfPlayers int
	correctInput bool
}

var playerTests = []playerTest{
	{"0", 0, true}, {"1", 1, true}, {"2", 2, true}, {"3", 3, true}, {"10", 10, false}, {"a", 99, false},
}

func TestUserCantGiveBadDataForNumOfPlayers(t *testing.T) {
	for _, pt := range playerTests {
		nop, ci := getPlayerCount(pt.input)

		if nop != pt.numOfPlayers {
			t.Errorf("Passing '"+pt.input+"' should result in %t as ready to play", pt.numOfPlayers)
		}

		if ci != pt.correctInput {
			t.Errorf("Passing '"+pt.input+"' should result in %t as the correct input", pt.correctInput)
		}
	}
}
