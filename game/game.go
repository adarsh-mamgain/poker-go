package game

import (
	"math/rand"
	"poker-go-cli/game/cards"
	"poker-go-cli/game/player"
	"time"
)

type GameState int

const (
	PreFlop GameState = iota
	Flop
	Turn
	River
	Showdown
)

type Game struct {
	Players     []*player.Player
	Deck        []cards.Card
	Community   []cards.Card
	Pot         int
	CurrentBet  int
	State       GameState
	DealerIndex int
}

func NewGame(players []*player.Player) *Game {
	return &Game{
		Players:     players,
		Deck:        cards.NewDeck(),
		Community:   make([]cards.Card, 0, 5),
		Pot:         0,
		CurrentBet:  0,
		State:       PreFlop,
		DealerIndex: 0,
	}
}

func (g *Game) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.Deck), func(i, j int) {
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	})
}

func (g *Game) DealCards() {
	g.ShuffleDeck()
	// Deal 2 cards to each player
	for i := 0; i < 2; i++ {
		for _, p := range g.Players {
			if len(g.Deck) > 0 {
				card := g.Deck[len(g.Deck)-1]
				g.Deck = g.Deck[:len(g.Deck)-1]
				p.AddCard(card)
			}
		}
	}
}

func (g *Game) DealFlop() {
	for i := 0; i < 3; i++ {
		if len(g.Deck) > 0 {
			card := g.Deck[len(g.Deck)-1]
			g.Deck = g.Deck[:len(g.Deck)-1]
			g.Community = append(g.Community, card)
		}
	}
	g.State = Flop
}

func (g *Game) DealTurn() {
	if len(g.Deck) > 0 {
		card := g.Deck[len(g.Deck)-1]
		g.Deck = g.Deck[:len(g.Deck)-1]
		g.Community = append(g.Community, card)
	}
	g.State = Turn
}

func (g *Game) DealRiver() {
	if len(g.Deck) > 0 {
		card := g.Deck[len(g.Deck)-1]
		g.Deck = g.Deck[:len(g.Deck)-1]
		g.Community = append(g.Community, card)
	}
	g.State = River
}

func (g *Game) ResetRound() {
	g.Deck = cards.NewDeck()
	g.Community = g.Community[:0]
	g.Pot = 0
	g.CurrentBet = 0
	g.State = PreFlop
	for _, p := range g.Players {
		p.ClearHand()
	}
} 