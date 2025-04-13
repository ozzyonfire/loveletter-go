package game

import (
	"errors"
	"fmt"
	"slices"
)

type Player struct {
	Name      string
	hand      []Card
	discards  []Card
	tokens    int
	protected bool // the handmaid can protect a player from being picked
}

func (p Player) String() string {
	if p.protected {
		return fmt.Sprint(p.Name, "(Protected)")
	}
	return p.Name
}

func (p *Player) PromptChoice() Card {
	index := Prompt(p.hand)
	return p.PlayCard(index)
}

func (p *Player) Discard(card Card) {
	p.discards = append(p.discards, card)
}

func (p *Player) PlayCard(index int) Card {
	cardToPlay := p.hand[index]
	p.hand = slices.Delete(p.hand, index, index)
	if _, ok := cardToPlay.(*Princess); ok {
		// the player is out
		fmt.Println("played the princess!")
	}
	p.Discard(cardToPlay)
	return cardToPlay
}

func (p *Player) DiscardHand() error {
	if len(p.hand) != 1 {
		return errors.New("player's hand has too many cards")
	}

	p.PlayCard(0)
	return nil
}

func (p *Player) Draw(game *Game) error {
	card, err := game.draw()
	if err != nil {
		return err
	}

	p.hand = append(p.hand, card)
	return nil
}
