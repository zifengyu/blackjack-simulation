package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDealingShoe(t *testing.T) {
	ds := NewDealingShoe(1)	
	ds.Shuffle()
	require.Equal(t, 52, ds.RemainCards(), "one deck should have 52 cards")

	ds = NewDealingShoe(2)	
	ds.Shuffle()
	require.Equal(t, 104, ds.RemainCards(), "two decks should have 104 cards")

	ds = NewDealingShoe(4)	
	ds.Shuffle()
	require.Equal(t, 208, ds.RemainCards(), "four decks should have 208 cards")
}

func TestDrawCard(t *testing.T) {
	cardCount1 := make([]int, 11)
	cardCount2 := make([]int, 11)
	var drawed1 []int
	var drawed2 []int

	ds := NewDealingShoe(2)

	ds.Shuffle()
	for ds.RemainCards() > 0 {
		card := ds.Draw()
		cardCount1[card]++
		drawed1 = append(drawed1, card)
	}

	ds.Shuffle()
	for ds.RemainCards() > 0 {
		card := ds.Draw()
		cardCount2[card]++
		drawed2 = append(drawed2, card)
	}

	for i := 1; i <= 9; i++ {
		require.Equal(t, 8, cardCount1[i])
		require.Equal(t, 8, cardCount2[i])
	}
	require.Equal(t, 32, cardCount1[10])
	require.Equal(t, 32, cardCount2[10])
	require.Equal(t, 0, cardCount1[0])
	require.Equal(t, 0, cardCount2[0])

	count := 0
	for i := 0; i < len(drawed1); i++ {
		if drawed1[i] != drawed2[i] {
			count++
		}
	}
	require.True(t, count > 0)
}
