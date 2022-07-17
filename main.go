package main

import (
	"fmt"
	AlphabetGame "project/game"
)

func main() {
	// Not interested in ui really
	// sample use case
	game := AlphabetGame.NewGame()

	game.AddPlayer("erfan")
	game.AddPlayer("emad")

	gameChar, err := game.Start()
	if err != nil {
		// do sth
	}
	fmt.Printf("The Alphabet for this round is %s", string(gameChar))

	winner, _ := game.End()

	fmt.Printf("Player %d won!", winner)
}
