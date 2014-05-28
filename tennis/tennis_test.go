package tennis_test

import (
	"github.com/wingyplus/kata/tennis"
	"testing"
)

func TestPlayer(t *testing.T) {
	var player *tennis.Player = tennis.NewPlayer()

	if player == nil {
		t.Errorf("player expected not nil")
	}
}

func TestTennisHasTwoPlayers(t *testing.T) {
	var player1 = tennis.NewPlayer()
	var player2 = tennis.NewPlayer()
	var tenniz *tennis.Tennis = tennis.NewTennis(player1, player2)

	if tenniz.PlayerOne() != player1 {
		t.Errorf("expected %v but was %v", tenniz.PlayerOne(), player1)
	}
	if tenniz.PlayerTwo() != player2 {
		t.Errorf("expected %v but was %v", tenniz.PlayerTwo(), player2)
	}
}

func TestPlayerShouldHasZeroScore(t *testing.T) {
	var player *tennis.Player = tennis.NewPlayer()

	if player.Score() != 0 {
		t.Errorf("expected player has score 0 but was %d", player.Score())
	}
}
