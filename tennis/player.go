package tennis

type Player struct {
	score int
}

func NewPlayer() *Player {
	return &Player{}
}

func (player *Player) Score() int {
	return player.score
}
