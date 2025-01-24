// card/card.go
package card

type Card uint8
type Suit uint8

// Rank index is the card number - 2. Eg a four is index 2.
type Rank uint8

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Suit returns the suit of the card
func (c Card) Suit() Suit {
	return Suit(c / 13)
}

// Rank returns the rank of the card
func (c Card) Rank() Rank {
	return Rank(c % 13)
}

// NewCard creates a card from a suit and rank
func NewCard(s Suit, r Rank) Card {
	return Card(int(s)*13 + int(r))
}

// IsValidCard checks if a card number is valid
func IsValidCard(c Card) bool {
	return c < 52
}

// String returns a human-readable representation of the card
func (c Card) String() string {
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := []string{"♠", "♥", "♦", "♣"}
	return ranks[c.Rank()] + suits[c.Suit()]
}
