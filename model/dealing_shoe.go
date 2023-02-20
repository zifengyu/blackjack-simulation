package model

import (
	"math/rand"
)

func Test() int {
	return 42
}

type DealingShoe struct {
	decks       int
	cards       []int
	cardsToDraw []int
}

func NewDealingShoe(decks int) *DealingShoe {
	ds := &DealingShoe{decks: decks}
	ds.Reload()
	return ds
}

func (ds *DealingShoe) Reload() {
	ds.cards = make([]int, 0, 52*ds.decks)
	for d := 0; d < ds.decks; d++ {
		for card := 1; card <= 13; card++ {
			if card < 10 {
				ds.cards = append(ds.cards, card, card, card, card)
			} else {
				ds.cards = append(ds.cards, 10, 10, 10, 10)
			}
		}
	}
}

func (ds *DealingShoe) Remove(card int) {
	for i, c := range ds.cards {
		if c == card {
			ds.cards[i] = ds.cards[len(ds.cards)-1]
			ds.cards = ds.cards[:len(ds.cards)-1]
			return
		}
	}
}

func (ds *DealingShoe) Shuffle() {
	ds.cardsToDraw = make([]int, len(ds.cards))
	copy(ds.cardsToDraw, ds.cards)
	rand.Shuffle(len(ds.cardsToDraw), func(i, j int) {
		ds.cardsToDraw[i], ds.cardsToDraw[j] = ds.cardsToDraw[j], ds.cardsToDraw[i]
	})
}

func (ds *DealingShoe) RemainCards() int {
	return len(ds.cardsToDraw)
}

func (ds *DealingShoe) Draw() int {
	if ds.RemainCards() == 0 {
		ds.Shuffle()
	}

	card := ds.cardsToDraw[0]
	ds.cardsToDraw = ds.cardsToDraw[1:]
	return card
}
