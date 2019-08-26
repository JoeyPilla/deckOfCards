package deck

import (
	"math/rand"
	"sort"
	"time"

	"../card"
)

type Deck struct {
	Deck []card.Card
}

func NewDeck() Deck {
	return Deck{
		card.NewDeck(),
	}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	deck := []card.Card{}
	for i := len(d.Deck); i > 0; i-- {
		rand := rand.Intn(i)
		deck = append(deck, d.Deck[rand])
		d.Deck[rand] = d.Deck[i-1]
	}
	d.Deck = deck
}

func (d Deck) Len() int {
	return len(d.Deck)
}

func (d Deck) Swap(i, j int) {
	d.Deck[i], d.Deck[j] = d.Deck[j], d.Deck[i]
}

func (d Deck) Less(i, j int) bool {
	if d.Deck[i].Suit == card.Joker {
		return false
	}
	if d.Deck[i].Suit == d.Deck[j].Suit {
		return d.Deck[i].Rank < d.Deck[j].Rank
	} else {
		return d.Deck[i].Suit < d.Deck[j].Suit
	}
}

func (d *Deck) Sort() {
	sort.Sort(d)
}

func (d *Deck) Filter(rank card.Rank) {
	cards := []card.Card{}
	for _, card := range d.Deck {
		if card.Rank != rank {
			cards = append(cards, card)
		}
	}
	d.Deck = cards
}

func (d *Deck) AddJokers(amount int) {
	for ; amount > 0; amount-- {
		d.Deck = append(d.Deck, card.Card{
			Suit: card.Joker,
		})
	}
}

func (d *Deck) AddDecks(amount int) {
	for ; amount > 0; amount-- {
		d.Deck = append(d.Deck, NewDeck().Deck...)
	}
}
