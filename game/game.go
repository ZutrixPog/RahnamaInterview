package game

import (
	"errors"
	"math/rand"
)

type State int

const (
	Stopped State = iota
	OnGoing
	Result
)

type Game struct {
	players       map[int]*Player
	state         State
	currentLetter rune
}

func NewGame() *Game {
	return &Game{
		players: make(map[int]*Player),
		state:   Stopped,
	}
}

func (g *Game) Start() (rune, error) {
	if len(g.players) == 0 && g.state != OnGoing {
		return 0, errors.New("There aren't enough players in the game!")
	}

	randomChar := 'a' + rune(rand.Intn(26))
	g.currentLetter = randomChar
	g.state = OnGoing
	return randomChar, nil
}

func (g *Game) StartWithCustomChar(char rune) error {
	if len(g.players) == 0 && g.state != OnGoing {
		return errors.New("There aren't enough players in the game!")
	}

	g.currentLetter = char
	g.state = OnGoing
	return nil
}

func (g *Game) AddPlayer(name string) (int, error) {
	if g.state == OnGoing {
		return 0, errors.New("You cant add a player to a ongoing game!")
	}
	playerId := len(g.players)
	g.players[playerId] = NewPlayer(name)
	return playerId, nil
}

func (g *Game) RemovePlayer(playerId int) error {
	if g.state == Stopped {
		return errors.New("You cant add a player to a ongoing game!")
	}
	delete(g.players, playerId)
	return nil
}

func (g *Game) AnswerForPlayer(playerId int, answer Answer) {
	g.players[playerId].setAnswer(answer)
}

func (g *Game) End() (int, error) {
	if g.state != OnGoing {
		return -1, errors.New("The game hasnt started yet!")
	}

	g.calculateScores()
	winner, _ := g.drawWinner()
	g.clear()
	return winner, nil
}

func (g *Game) drawWinner() (int, int) {
	winner := -1
	maxScore := 0
	for id, player := range g.players {
		if player.Score > maxScore {
			maxScore = player.Score
			winner = id
		}
	}
	return winner, maxScore
}

func (g *Game) calculateScores() {
	// Now that i think i didnt implement an efficient data structure
	// pick one player
	for id, player := range g.players {
		// check for each of their answers
		for i, ans := range player.answers {
			// zero score if it doesnt start with the correct letter
			if ans == "" {
				continue
			}
			if []rune(ans)[0] != g.currentLetter {
				continue
			}
			// check the answer with other players answer and see if it's a duplication
			for id2, player2 := range g.players {
				if id == id2 {
					continue
				}
				// check for duplication
				if ans == player2.answers[i] {
					player.Score += 5
					break
				}
			}
			// everything okay? then give them full score
			player.Score += 10
		}
	}
}

func (g *Game) clear() {
	for _, player := range g.players {
		player.ClearAnswer()
	}
}
