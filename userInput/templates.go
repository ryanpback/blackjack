package userInput

// let's standardize the verbiage. Also, with the upper/lower syntax, we can let
// the user just press enter and it'd be the same as entering the uppercase
// character option
var Intro = `
Ready to play BlackJack?
========================
Yes (Y) | Quit (q):
`

var PlayerIsHuman = `
Is this player human? (Y|n)
`

var Outro = `
Goodbye
`

var AskNumberOfPlayers = `
How many players including you and the dealer? ({{ .Min }} - {{ .Max }}) :
`

var InvalidOption = `Sorry, that is not a valid option.`

var GameStart = `
Let's play!
===========
`

var dealerCreated = `Dealer. They are controlled by the computer.`

var getPlayerName = `What is this player's name? :`
