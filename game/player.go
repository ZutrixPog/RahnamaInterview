package game

type Player struct {
	Name    string
	Score   int
	answers []string // for ease of calculation
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:    name,
		Score:   0,
		answers: make([]string, 0),
	}
}

func (p *Player) setAnswer(answer Answer) error {
	// do your checks and return error here (NO checks right now)
	for _, field := range []interface{}{answer.Name, answer.City, answer.Color, answer.Food} {
		p.answers = append(p.answers, field.(string))
	}
	return nil
}

func (p *Player) ClearAnswer() {
	p.answers = nil
	p.Score = 0
}

type Answer struct {
	Name  string
	City  string
	Color string
	Food  string
}
