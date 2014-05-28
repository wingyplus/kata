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
