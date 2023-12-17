package entities

import "math/rand"

var scoreChoices = []int{0, 2, 3}
var playerChoices = []int{0, 1, 2, 3, 4}

func NewTeam(id, name string, players []Player) ITeam {
	return &Team{Id: id,
		Name:    name,
		Players: players,
		Points:  0}
}

func (t *Team) Score() {
	point := randomScorePoints()

	if point != 0 {
		scorerOrder, assistOrder := randomPickScorerAssist()
		t.Points += point
		t.Players[scorerOrder].Points += point
		t.Players[assistOrder].Assists += 1
	}
}

func (t *Team) GetPoints() int {
	return t.Points
}

func randomScorePoints() int {
	i := rand.Intn(len(scoreChoices))
	return scoreChoices[i]
}

func randomPickScorerAssist() (int, int) {
	scorer := rand.Intn(5)
	restPlayerChoices := append(playerChoices[:scorer], playerChoices[scorer+1:]...)
	assist := rand.Intn(len(restPlayerChoices))
	return scorer, assist
}
