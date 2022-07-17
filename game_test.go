package main

import (
	"fmt"
	"project/game"
	"testing"
)

func TestGame(t *testing.T) {
	g := game.NewGame()

	erfanId, _ := g.AddPlayer("erfan")
	emadId, _ := g.AddPlayer("emad")
	aliId, _ := g.AddPlayer("ali")

	err := g.StartWithCustomChar('a')
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Games Alphabet is %s\n", "a")

	g.AnswerForPlayer(erfanId, game.Answer{Name: "ali", City: "abudabi", Color: "", Food: "adasi"})
	g.AnswerForPlayer(emadId, game.Answer{Name: "ali", City: "", Color: "", Food: ""})
	g.AnswerForPlayer(aliId, game.Answer{Name: "akbar", City: "acity", Color: "acolor", Food: "afood"})

	winner, err := g.End()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("your winner is player %d\n", winner)
	if winner != aliId {
		t.Errorf("Wrong output for winner")
	}
}
