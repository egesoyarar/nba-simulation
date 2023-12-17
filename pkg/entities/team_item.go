package entities

type Team struct {
	Id      string
	Name    string
	Players []Player
	Points  int
}

type ITeam interface {
	Score()
}
