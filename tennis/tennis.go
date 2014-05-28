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

type Tennis struct {
	player1, player2 *Player
}

func NewTennis(player1, player2 *Player) *Tennis {
	return &Tennis{player1, player2}
}

func (tenniz *Tennis) PlayerOne() *Player {
	return tenniz.player1
}

func (tenniz *Tennis) PlayerTwo() *Player {
	return tenniz.player2
}
