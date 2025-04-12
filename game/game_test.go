package game

import (
	"fmt"
	"testing"
)

func TestSwapHands(t *testing.T) {
	game, err := NewGame(3)
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	player1, player2 := &game.players[0], &game.players[1]

	// Give players test hands
	king := NewKing()
	player1.hand = []Card{king}
	player2.hand = []Card{NewCard("Guard", 1)}

	fmt.Println(player1.hand, player2.hand)
	fmt.Println("Swap")
	game.SwapHands(player1, player2)
	fmt.Println(player1.hand, player2.hand)

	if player2.hand[0] != king {
		t.Error("Swap failed.")
	}
}
