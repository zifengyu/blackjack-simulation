package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandCards(t *testing.T) {
	dealer := NewDealer()
	require.Equal(t, 0, dealer.Value(), "hand value should be 0 before get any card")

	dealer.GetCard(4)
	require.Equal(t, 4, dealer.Value())

	dealer.CleanCards()
	require.Equal(t, 0, dealer.Value())

	dealer.GetCard(8)
	require.Equal(t, 8, dealer.Value())
	require.False(t, dealer.IsBust())
	require.False(t, dealer.IsStand())

	dealer.GetCard(8)
	require.Equal(t, 16, dealer.Value())
	require.False(t, dealer.IsBust())
	require.False(t, dealer.IsStand())

	dealer.GetCard(1)
	require.Equal(t, 17, dealer.Value())
	require.False(t, dealer.IsBust())
	require.True(t, dealer.IsStand())

	dealer.GetCard(1)
	require.Equal(t, 18, dealer.Value())
	require.False(t, dealer.IsBust())
	require.True(t, dealer.IsStand())

	dealer.GetCard(10)
	require.Equal(t, 28, dealer.Value())
	require.True(t, dealer.IsBust())

	dealer.CleanCards()
	dealer.GetCard(1)
	dealer.GetCard(1)
	require.Equal(t, 12, dealer.Value())
	require.True(t, dealer.IsSoft())
	require.False(t, dealer.IsBust())
	require.False(t, dealer.IsStand())

	dealer.GetCard(10)
	require.Equal(t, 12, dealer.Value())
	require.False(t, dealer.IsSoft())
	require.False(t, dealer.IsBust())
	require.False(t, dealer.IsStand())

	dealer.GetCard(1)
	require.Equal(t, 13, dealer.Value())
	require.False(t, dealer.IsSoft())
	require.False(t, dealer.IsBust())
	require.False(t, dealer.IsStand())

	dealer.GetCard(9)
	require.Equal(t, 22, dealer.Value())
	require.False(t, dealer.IsSoft())
	require.True(t, dealer.IsBust())

	dealer.CleanCards()
	require.False(t, dealer.CanSplit())
	dealer.GetCard(8)
	require.False(t, dealer.CanSplit())
	dealer.GetCard(8)
	require.True(t, dealer.CanSplit())
	dealer.GetCard(8)
	require.False(t, dealer.CanSplit())

	dealer.CleanCards()
	require.False(t, dealer.CanSplit())
	dealer.GetCard(8)
	dealer.GetCard(7)
	require.False(t, dealer.CanSplit())
}

func TestBasicStandPlayer(t *testing.T) {
	p := NewBasicStandPlayer(5)
	p.GetCard(3)
	require.False(t, p.IsStand())
	p.GetCard(2)
	require.True(t, p.IsStand())
	p.GetCard(2)
	require.True(t, p.IsStand())
}
