package game

import (
	"errors"
	"fmt"
)

type Card interface {
	Effect(game *Game, player *Player) error
	String() string
}

type BaseCard struct {
	Name  string
	Value int
}

func NewCard(name string, value int) *BaseCard {
	return &BaseCard{
		Name:  name,
		Value: value,
	}
}

func (b *BaseCard) String() string {
	return fmt.Sprint(b.Name, " ", b.Value)
}

func (b *BaseCard) Effect(game *Game, player *Player) error {
	fmt.Println("Playing", b)
	return errors.New("you can't play a BaseCard")
}

// Princess
type Princess struct {
	BaseCard
}

func (p *Princess) Effect(game *Game, player *Player) error {
	// Do nothing
	return nil
}

// Countess
type Countess struct {
	BaseCard
}

func (c *Countess) Effect(game *Game, player *Player) error {
	// No effect after card is played
	return nil
}

// King
type King struct {
	BaseCard
}

func NewKing() *King {
	return &King{
		BaseCard: *NewCard("King", 7),
	}
}

func (k *King) choosePlayer(game *Game, player *Player) *Player {
	// Pick a player and switch hands
	fmt.Println("Pick a player to swap hands with:")
	index := 1

	otherPlayers := make([]*Player, 0, len(game.players)-1)

	for _, p := range game.players {
		if p.Name != player.Name {
			fmt.Printf("%d) %v\n", index, p.Name)
			otherPlayers = append(otherPlayers, &p)
			index++
		}
	}

	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Invalid choice.")
		return k.choosePlayer(game, player)
	}

	if input < 1 || input > len(otherPlayers) {
		fmt.Println("Invalid choice. Please select a valid player number.")
		return k.choosePlayer(game, player) // retry
	}

	return otherPlayers[input-1]
}

func (k *King) Effect(game *Game, player *Player) error {

	if len(game.players) == 1 {
		return errors.New("only one player left")
	}

	chosenPlayer := k.choosePlayer(game, player)
	game.SwapHands(player, chosenPlayer)

	return nil
}

// Chancellor

type Chancellor struct {
	BaseCard
}

func NewChancellor() *Chancellor {
	return &Chancellor{
		BaseCard: *NewCard("Chancellor", 6),
	}
}

func (c *Chancellor) Effect(game *Game, player *Player) error {
	cardsDrawn := make([]Card, 0, 2)

	card1, err1 := game.draw()
	if err1 != nil {
		if _, ok := err1.(*DeckEmptyError); !ok {
			return err1
		}
	} else {
		cardsDrawn = append(cardsDrawn, card1)
	}

	card2, err2 := game.draw()
	if err2 != nil {
		if _, ok := err2.(*DeckEmptyError); !ok {
			return err2
		}
	} else {
		cardsDrawn = append(cardsDrawn, card2)
	}

	hand := append(player.hand, cardsDrawn...)
	index := Prompt(hand)

	// Keep the selected card
	player.hand = []Card{hand[index]}

	// Put remaining cards at bottom of deck
	for i, card := range hand {
		if i != index {
			game.deck = append([]Card{card}, game.deck...)
		}
	}

	return nil
}

// Prince

type Prince struct {
	BaseCard
}

func NewPrince() *Prince {
	return &Prince{
		BaseCard: *NewCard("Prince", 5),
	}
}

func (p *Prince) Effect(game *Game, player *Player) error {
	index := Prompt(game.players)
	var chosenUser = game.players[index]
	chosenUser.DiscardHand()
	return nil
}

// Handmaid

type Handmaid struct {
	BaseCard
}

func NewHandmaid() *Handmaid {
	return &Handmaid{
		BaseCard: *NewCard("Handmaid", 4),
	}
}

func (h *Handmaid) Effect(game *Game, player *Player) error {
	player.protected = true
	return nil
}
