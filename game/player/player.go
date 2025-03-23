package player

import (
	"poker-go-cli/game/cards"
)

type Player struct {
	Name     string
	Chips    int
	Hand     []cards.Card
	IsActive bool
	IsDealer bool
}

func NewPlayer(name string, chips int) *Player {
	return &Player{
		Name:     name,
		Chips:    chips,
		Hand:     make([]cards.Card, 0, 5),
		IsActive: true,
		IsDealer: false,
	}
}

func (p *Player) AddCard(card cards.Card) {
	p.Hand = append(p.Hand, card)
}

func (p *Player) ClearHand() {
	p.Hand = p.Hand[:0]
}

func (p *Player) PlaceBet(amount int) bool {
	if amount > p.Chips {
		return false
	}
	p.Chips -= amount
	return true
}

func (p *Player) AddChips(amount int) {
	p.Chips += amount
} 