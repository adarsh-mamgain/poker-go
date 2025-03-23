package cards

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

type Rank int

const (
	Two Rank = iota + 2
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

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := []string{"♥", "♦", "♣", "♠"}
	return ranks[c.Rank-2] + suits[c.Suit]
}

func NewDeck() []Card {
	deck := make([]Card, 52)
	index := 0
	for suit := Hearts; suit <= Spades; suit++ {
		for rank := Two; rank <= Ace; rank++ {
			deck[index] = Card{Suit: suit, Rank: rank}
			index++
		}
	}
	return deck
} 