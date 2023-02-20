package main

import (
	"fmt"
	"time"

	"github.com/zifengyu/blackjack-simulation/model"
)

const SIM_ROUND = 1000000

func simulateMimicDealer() float64 {
	ds := model.NewDealingShoe(1)
	dealer := model.NewDealer()
	player := model.NewDealer()
	return model.SimBlackjack(ds, dealer, player, SIM_ROUND)
}

func simulateNeverBust() float64 {
	ds := model.NewDealingShoe(1)
	dealer := model.NewDealer()
	player := model.NewNeverBustPlayer()
	return model.SimBlackjack(ds, dealer, player, SIM_ROUND)
}

func main() {
	start := time.Now()
	fmt.Println()
	fmt.Println("Mimic dealer:", simulateMimicDealer())
	fmt.Println("Never bust:  ", simulateNeverBust())

	elapsed := time.Since(start)
	fmt.Printf("Finished in %s\n", elapsed)
}
