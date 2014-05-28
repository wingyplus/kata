package tennis

type Player struct{}

func NewPlayer() *Player {
	return &Player{}
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
