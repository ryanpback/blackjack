package player

import "testing"

func TestCanSetAPlayersName(t *testing.T) {
	p := Player{}
	if p.Name != "" {
		t.Errorf("Default player name should be blank")
	}

	p.SetName("Jimmy Smith")
	if p.Name != "Jimmy Smith" {
		t.Errorf("Player name not set correctly")
	}
}

func TestPlayerHasAShortDisplayName(t *testing.T) {
	p := Player{
		Name: "Kurtis Holsapple",
	}

	if p.ShortName() != "KH" {
		t.Errorf("The short name was not generated correctly. Expected %s, got %s", "KH", p.ShortName())
	}
}
