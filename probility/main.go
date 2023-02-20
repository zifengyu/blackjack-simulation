package main

import (
	"fmt"
	"time"

	"github.com/zifengyu/blackjack-simulation/model"
)

func blackjackProbability(deck int, n int) float64 {
	blackjackCount := 0
	ds := model.NewDealingShoe(deck)
	dealer := model.NewDealer()

	for i := 0; i < n; i++ {
		// ds.Shuffle()
		dealer.CleanCards()
		dealer.GetCard(ds.Draw())
		dealer.GetCard(ds.Draw())
		if dealer.Value() == 21 {
			blackjackCount++
		}
	}

	return float64(blackjackCount) / float64(n)
}

func main() {
	start := time.Now()

	fmt.Println("Blackjack probability for 1 deck is", blackjackProbability(1, 10000000))
	fmt.Println("Blackjack probability for 6 deck is", blackjackProbability(6, 10000000))

	elapsed := time.Since(start)
	fmt.Printf("\nFinished in %s\n", elapsed)
}
