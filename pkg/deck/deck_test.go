package deck_test

import (
	"testing"

	"github.com/dgunzy/card/pkg/card"
	"github.com/dgunzy/card/pkg/deck"
)

func TestNew(t *testing.T) {
	d := deck.New()
	if d.RemainingCardsCount() != 52 {
		t.Errorf("expected 52 cards, got %d", d.RemainingCardsCount())
	}
}

func TestDraw(t *testing.T) {
	d := deck.New()
	cr, err := d.Draw()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if d.RemainingCardsCount() != 51 {
		t.Errorf("expected 51 cards, got %d", d.RemainingCardsCount())
	}
	if !card.IsValidCard(cr) {
		t.Errorf("card value %d is not valid", cr)
	}
}

func TestDrawMany(t *testing.T) {
	d := deck.New()

	// Test successful draw
	cards, err := d.DrawMany(5)
	if err != nil {
		t.Errorf("unexpected error drawing 5 cards: %v", err)
	}
	if len(cards) != 5 {
		t.Errorf("expected 5 cards, got %d", len(cards))
	}
	if d.RemainingCardsCount() != 47 {
		t.Errorf("expected 47 cards remaining, got %d", d.RemainingCardsCount())
	}

	// Test drawing too many cards
	_, err = d.DrawMany(48)
	if err == nil {
		t.Error("expected error drawing too many cards, got nil")
	}
}

func TestPeekTop(t *testing.T) {
	d := deck.New()

	// Test peek on full deck
	topCard, err := d.PeekTop()
	if err != nil {
		t.Errorf("unexpected error peeking top card: %v", err)
	}

	// Draw a card and verify it was the one we peeked
	drawnCard, err := d.Draw()
	if err != nil {
		t.Errorf("unexpected error drawing card: %v", err)
	}
	if topCard != drawnCard {
		t.Errorf("peeked card %v different from drawn card %v", topCard, drawnCard)
	}

	// Empty deck and test peek
	for i := 0; i < 51; i++ {
		_, _ = d.Draw()
	}
	_, err = d.PeekTop()
	if err == nil {
		t.Error("expected error peeking empty deck, got nil")
	}
}

func TestInsertCards(t *testing.T) {
	d := deck.New()
	testCard := card.NewCard(card.Spades, card.Ace)

	// Test InsertTop
	d.InsertTop(testCard)
	if d.RemainingCardsCount() != 53 {
		t.Errorf("expected 53 cards after insert, got %d", d.RemainingCardsCount())
	}
	topCard, _ := d.PeekTop()
	if topCard != testCard {
		t.Errorf("expected top card to be %v, got %v", testCard, topCard)
	}

	// Test InsertBottom
	d.Reset()
	d.InsertBottom(testCard)
	if d.RemainingCardsCount() != 53 {
		t.Errorf("expected 53 cards after insert, got %d", d.RemainingCardsCount())
	}
	// Draw all but last card
	for i := 0; i < 52; i++ {
		_, _ = d.Draw()
	}
	lastCard, _ := d.Draw()
	if lastCard != testCard {
		t.Errorf("expected bottom card to be %v, got %v", testCard, lastCard)
	}
}

func TestCardProperties(t *testing.T) {
	// Test card creation and properties
	testCases := []struct {
		suit     card.Suit
		rank     card.Rank
		expected card.Card
	}{
		{card.Spades, card.Ace, 12},     // (0 * 13) + 12 = 12
		{card.Hearts, card.King, 24},    // (1 * 13) + 11 = 24
		{card.Diamonds, card.Queen, 36}, // (2 * 13) + 10 = 36
		{card.Clubs, card.Ten, 47},      // (3 * 13) + 8 = 47
	}

	for _, tc := range testCases {
		c := card.NewCard(tc.suit, tc.rank)
		if uint8(c) != uint8(tc.expected) {
			t.Errorf("NewCard(%d, %d) = %d, want %d",
				uint8(tc.suit), uint8(tc.rank), uint8(c), uint8(tc.expected))
		}
		if c.Suit() != tc.suit {
			t.Errorf("card.Suit() = %d, want %d", uint8(c.Suit()), uint8(tc.suit))
		}
		if c.Rank() != tc.rank {
			t.Errorf("card.Rank() = %d, want %d", uint8(c.Rank()), uint8(tc.rank))
		}
	}
}

func TestCardString(t *testing.T) {
	testCases := []struct {
		card     card.Card
		expected string
	}{
		{card.NewCard(card.Spades, card.Ace), "A♠"},
		{card.NewCard(card.Hearts, card.King), "K♥"},
		{card.NewCard(card.Diamonds, card.Queen), "Q♦"},
		{card.NewCard(card.Clubs, card.Ten), "10♣"},
	}

	for _, tc := range testCases {
		if str := tc.card.String(); str != tc.expected {
			t.Errorf("card.String() = %v, want %v", str, tc.expected)
		}
	}
}

func TestDrawEmpty(t *testing.T) {
	d := deck.New()
	for i := 0; i < 52; i++ {
		_, err := d.Draw()
		if err != nil {
			t.Errorf("unexpected error drawing card %d: %v", i, err)
		}
	}
	if _, err := d.Draw(); err == nil {
		t.Error("expected error drawing from empty deck, got nil")
	}
}

func TestShuffle(t *testing.T) {
	d := deck.New()
	originalOrder := make([]card.Card, 52)
	for i := 0; i < 52; i++ {
		card, err := d.Draw()
		if err != nil {
			t.Fatalf("failed to draw card: %v", err)
		}
		originalOrder[i] = card
	}

	d.Reset()
	d.Shuffle()

	matchingCards := 0
	for i := 0; i < 52; i++ {
		card, err := d.Draw()
		if err != nil {
			t.Fatalf("failed to draw card: %v", err)
		}
		if card == originalOrder[i] {
			matchingCards++
		}
	}

	if matchingCards > 45 {
		t.Errorf("shuffle appears ineffective: %d/52 cards in same position", matchingCards)
	}
}

func TestReset(t *testing.T) {
	d := deck.New()
	for i := 0; i < 10; i++ {
		_, err := d.Draw()
		if err != nil {
			t.Errorf("unexpected error drawing card %d: %v", i, err)
		}
	}
	d.Reset()
	if d.RemainingCardsCount() != 52 {
		t.Errorf("expected 52 cards after reset, got %d", d.RemainingCardsCount())
	}
}
