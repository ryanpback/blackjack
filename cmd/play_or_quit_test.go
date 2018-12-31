package main

import (
	"testing"
)

type playQuitTest struct {
	input        string
	readyToPlay  bool
	correctInput bool
}

var playQuitTests = []playQuitTest{
	{"Y", true, true}, {"Q", false, true}, {"L", false, false}, {"1", false, false},
}

func TestUserCantGiveBadDataForWelcomeScreen(t *testing.T) {
	for _, pqt := range playQuitTests {
		rtp, ci := playOrQuit(pqt.input)

		if rtp != pqt.readyToPlay {
			t.Errorf("Passing '"+pqt.input+"' should result in %t as ready to play", pqt.readyToPlay)
		}

		if ci != pqt.correctInput {
			t.Errorf("Passing '"+pqt.input+"' should result in %t as the correct input", pqt.correctInput)
		}
	}
}
