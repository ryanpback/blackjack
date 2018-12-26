package main

import (
	"testing"
)

type welcomeTest struct {
	input        string
	readyToPlay  bool
	correctInput bool
}

var welcomeTests = []welcomeTest{
	{"Y", true, true}, {"Q", false, true}, {"L", false, false}, {"1", false, false},
}

func TestUserCantGiveBadData(t *testing.T) {
	for _, wt := range welcomeTests {
		rtp, ci := playOrQuit(wt.input)

		if rtp != wt.readyToPlay {
			t.Errorf("Passing '"+wt.input+"' should result in %t as ready to play", wt.readyToPlay)
		}

		if ci != wt.correctInput {
			t.Errorf("Passing '"+wt.input+"' should result in %t as the correct input", wt.readyToPlay)
		}
	}
}
