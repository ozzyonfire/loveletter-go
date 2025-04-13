package main

import "github.com/ozzyonfire/go-loveletter/game"

// This is the main function
func main() {
	game1, err := game.NewGame(4)
	if err != nil {
		panic(err)
	}

	game1.StartRound()

	game1.Print()
}
