package game

import (
	"slices"
)

type Player struct {
	Name     string
	hand     []Card
	discards []Card
	tokens   int
}

func (p *Player) PromptChoice() Card {
	index := Prompt(p.hand)
	cardToPlay := p.hand[index]
	p.hand = slices.Delete(p.hand, index, index)
	p.Discard(cardToPlay)
	return cardToPlay
}

func (p *Player) Discard(card Card) {
	p.discards = append(p.discards, card)
}

func (p *Player) Draw(game *Game) error {
	card, err := game.draw()
	if err != nil {
		return err
	}

	p.hand = append(p.hand, card)
	return nil
}
