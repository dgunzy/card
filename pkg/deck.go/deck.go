package deck

import (
	"errors"
	"math/rand"

	"github.com/dgunzy/card/pkg/card"
)

// Deck represents a deck of Cards.
type Deck struct {
	Cards []card.Card
}

// New creates a new deck of Cards.
func New() *Deck {
	d := &Deck{}
	d.Reset()
	return d
}

// Reset resets the deck to its original state.
func (d *Deck) Reset() {
	d.Cards = make([]card.Card, 52)
	for i := range d.Cards {
		d.Cards[i] = card.Card(i)
	}
}

// Draw draws a Card from the deck.
func (d *Deck) Draw() (card.Card, error) {
	if len(d.Cards) == 0 {
		return card.Card(0), errors.New("no Cards left in deck")
	}
	drawnCard := d.Cards[0]
	d.Cards = d.Cards[1:]
	return drawnCard, nil
}

// DrawMany draws n cards from the deck
func (d *Deck) DrawMany(n int) ([]card.Card, error) {
	if n > len(d.Cards) {
		return nil, errors.New("not enough cards in deck")
	}
	cards := make([]card.Card, n)
	for i := 0; i < n; i++ {
		card, err := d.Draw()
		if err != nil {
			return nil, err
		}
		cards[i] = card
	}
	return cards, nil
}

// Shuffle shuffles the deck.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// RemainingCardsCount returns the number of Cards left in the deck.
func (d *Deck) RemainingCardsCount() int {
	return len(d.Cards)
}

// PeekTop returns the top card without removing it
func (d *Deck) PeekTop() (card.Card, error) {
	if len(d.Cards) == 0 {
		return card.Card(0), errors.New("no cards left in deck")
	}
	return d.Cards[0], nil
}

// InsertBottom puts a card at the bottom of the deck
func (d *Deck) InsertBottom(c card.Card) {
	d.Cards = append(d.Cards, c)
}

// InsertTop puts a card at the top of the deck
func (d *Deck) InsertTop(c card.Card) {
	d.Cards = append([]card.Card{c}, d.Cards...)
}
