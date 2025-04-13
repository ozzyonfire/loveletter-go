package game

import (
	"testing"
)

func TestKing(t *testing.T) {
	game, err := NewGame(3)
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	// Setup test players
	player2 := &game.players[1]
	player1 := &game.players[0]

	// Give players test hands
	king := NewKing()
	player1.hand = []Card{king}
	player2.hand = []Card{NewCard("Guard", 1)}

	// king.Effect(game, player1)
}
