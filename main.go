package main

import "github.com/ozzyonfire/go-loveletter/game"

func main() {
	game1, err := game.NewGame(4)
	if err != nil {
		panic(err)
	}

	game1.StartRound()

	game1.Print()
}
