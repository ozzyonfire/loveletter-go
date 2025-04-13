package game

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

type Game struct {
	players            []Player
	round              int
	deck               []Card
	topCard            Card
	currentPlayerIndex int
}

func (g *Game) Setup() {
	/* for a 2 to 8 player game we need the following:
	- Princess
	- Countess
	- King
	- Prince (x2)
	- Handmaid (x2)
	- Baron (x2)
	- Priest (x2)
	- Guard (x6)
	- Spy (x2)
	*/

	princess := NewCard("Princess", 9)      // If you play this, you are out
	countess := NewCard("Countess", 8)      // Must play if King or Prince in hand
	king := NewCard("King", 7)              // Choose a player and trade hands
	chancellor := NewCard("Chancellor", 6)  // Draw 2 cards from the deck. Choose 1, place the other 2 at the bottom of the deck in any order
	chancellor2 := NewCard("Chancellor", 6) // Draw 2 cards from the deck. Choose 1, place the other 2 at the bottom of the deck in any order
	prince := NewCard("Prince", 5)          // Choose any player, that player discards their hand (no effect) and draws a new one
	prince2 := NewCard("Prince", 5)         // Choose any player, that player discards their hand (no effect) and draws a new one
	handmaid := NewCard("Handmaid", 4)      // Protected for one turn
	handmaid2 := NewCard("Handmaid", 4)     // Protected for one turn
	baron := NewCard("Baron", 3)            // Compare with another player, the lowest value is out
	baron2 := NewCard("Baron", 3)           // Compare with another player, the lowest value is out
	priest := NewCard("Priest", 2)          // Look at another player's hand
	priest2 := NewCard("Priest", 2)         // Look at another player's hand
	guard := NewCard("Guard", 1)            // Choose a player and name a card (other than guard), if it matches, they are out.
	guard2 := NewCard("Guard", 1)           // Choose a player and name a card (other than guard), if it matches, they are out.
	guard3 := NewCard("Guard", 1)           // Choose a player and name a card (other than guard), if it matches, they are out.
	guard4 := NewCard("Guard", 1)           // Choose a player and name a card (other than guard), if it matches, they are out.
	guard5 := NewCard("Guard", 1)           // Choose a player and name a card (other than guard), if it matches, they are out.
	guard6 := NewCard("Guard", 1)           // Choose a player and name a card (other than guard), if it matches, they are out.
	spy := NewCard("Spy", 0)                // If discarded that round, get an extra token.
	spy2 := NewCard("Spy", 0)               // If discarded that round, get an extra token.

	deck := []Card{
		princess,
		countess,
		king,
		chancellor,
		chancellor2,
		prince,
		prince2,
		handmaid,
		handmaid2,
		baron,
		baron2,
		priest,
		priest2,
		guard,
		guard2,
		guard3,
		guard4,
		guard5,
		guard6,
		spy,
		spy2,
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	g.deck = deck
}

func NewGame(numOfPlayers int) (*Game, error) {
	if numOfPlayers <= 0 {
		return nil, errors.New("at least 2 players needed")
	}

	players := make([]Player, numOfPlayers)

	for i := 0; i < numOfPlayers; i++ {
		players[i] = Player{
			Name:     fmt.Sprint("Player", i+1),
			hand:     make([]Card, 0, 2),
			discards: make([]Card, 0, 5),
			tokens:   0,
		}
	}

	return &Game{
		players: players,
		round:   1,
	}, nil
}

func (g *Game) Print() {
	for i, player := range g.players {
		var currentPlayerIndicator = " "
		if g.currentPlayerIndex == i {
			currentPlayerIndicator = ">"
		}
		fmt.Println(currentPlayerIndicator, player.Name, player.hand)
	}
}

type DeckEmptyError struct{}

func (e *DeckEmptyError) Error() string {
	return "no cards left in deck"
}

func (g *Game) draw() (Card, error) {
	lastIdx := len(g.deck) - 1
	if lastIdx < 0 {
		return nil, &DeckEmptyError{}
	}
	card := g.deck[lastIdx]
	g.deck = g.deck[:lastIdx]
	return card, nil
}

func (g *Game) StartRound() error {
	g.Setup() // shuffle the deck

	// store the top card face down
	card, err := g.draw()
	if err != nil {
		return err
	}
	g.topCard = card

	if err := g.deal(); err != nil {
		return err
	}

	g.currentPlayerIndex = randRange(0, len(g.players))
	g.roundLoop()
	return nil
}

func (g *Game) deal() error {
	// deal out cards to players
	for i := range g.players {
		card, err := g.draw()
		if err != nil {
			return err
		}

		g.players[i].hand = append(g.players[i].hand, card)
	}
	return nil
}

func randRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (g *Game) getCurrentPlayer() *Player {
	return &g.players[g.currentPlayerIndex]
}

func (g *Game) roundLoop() {
	for {
		currentPlayer := g.getCurrentPlayer()

		// draw a card
		currentPlayer.Draw(g)

		// wait for player to play a card
		cardToPlay := currentPlayer.PromptChoice()

		cardToPlay.Effect(g, currentPlayer)

		// get the next player
		g.nextPlayer()
	}
}

func (g *Game) nextPlayer() {
	nextIndex := g.currentPlayerIndex + 1
	if nextIndex >= len(g.players) {
		nextIndex = 0
	}
	g.currentPlayerIndex = nextIndex
}

func (g *Game) SwapHands(player1 *Player, player2 *Player) {
	player1.hand, player2.hand = player2.hand, player1.hand
}

// TODO: players needs to be different from Round Players
func (g *Game) RemovePlayer(player Player) error {
	var index = -1
	for i, p := range g.players {
		if p.Name == player.Name {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("player not found")
	}

	g.players = slices.Delete(g.players, index, index+1)

	if index < g.currentPlayerIndex {
		g.currentPlayerIndex--
	}
	if g.currentPlayerIndex >= len(g.players) {
		g.currentPlayerIndex = 0
	}
	return nil
}
