package player

import (
	"strings"

	"github.com/lapubell/cards"
)

// Player this is a person who is playing the game
type Player struct {
	IsHuman bool // if false, the computer will automatically act on behalf of this Player
	Name    string
	Hand    []cards.Card
	Wallet  int // used to show how much money (in pennies) this player has
}

// SetName Set the name of the Player
func (p *Player) SetName(name string) {
	p.Name = name
}

// ShortName return a short version of the Player's name for shorter display
func (p *Player) shortName() string {
	if p.Name == "" {
		return ""
	}
	output := ""
	words := strings.Split(p.Name, " ")
	for _, word := range words {
		output = output + strings.ToUpper(string(word[0]))
	}
	return output
}

func (p Player) safeName() string {
	name := p.shortName()
	switch len(name) {
	case 0:
		return "     "
	case 1:
		return "  " + name + "  "
	case 2:
		return " " + name + "  "
	case 3:
		return " " + name + " "
	case 4:
		return " " + name
	case 5:
		return name
	default:
		return name[0:5]
	}
}

// Draw show a little character for this user
func (p Player) Draw() string {
	return `  o  
/| |\
` + p.safeName()
}

// SetupPlayer get this player ready to go
func SetupPlayer(name string, isHuman bool, wallet int) Player {
	p := Player{}
	p.SetName(name)
	p.IsHuman = isHuman
	p.Wallet = wallet

	return p
}
