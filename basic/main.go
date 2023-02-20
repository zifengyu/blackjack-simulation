package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/zifengyu/blackjack-simulation/model"
)

const SIM_ROUND = 1000000

func simluateBasicHard(dealerUp int, stand int) float64 {
	ds := model.NewDealingShoe(6)
	dealer := model.NewDealer()
	player := model.NewBasicStandPlayer(stand)
	return model.SimBasicStandHard(ds, dealer, player, dealerUp, SIM_ROUND)
}

func simluateBasicSoft(dealerUp int, stand int) float64 {
	ds := model.NewDealingShoe(6)
	dealer := model.NewDealer()
	player := model.NewBasicSoftStandPlayer(stand)
	return model.SimBasicStandSoft(ds, dealer, player, dealerUp, SIM_ROUND)
}

func simulateBasicV2(dealerUp int, hard int, soft int) float64 {
	ds := model.NewDealingShoe(1)
	dealer := model.NewDealer()
	player := model.NewBasicPlayer(hard, soft)
	return model.SimBasicStand(ds, dealer, player, dealerUp, SIM_ROUND)
}

func simulatePerfectBasicPlayer() float64 {
	ds := model.NewDealingShoe(1)
	dealer := model.NewDealer()
	player := model.NewPerfectBasicPlayer()
	return model.SimBlackjack(ds, dealer, player, SIM_ROUND)
}

func simulatePerfectBasicPlayerBiasedShoe() {
	for card := 1; card <= 10; card++ {
		ds := model.NewDealingShoe(1)
		n := 4
		if card == 10 {
			n = 16
		}
		for i := 0; i < n; i++ {
			ds.Remove(card)
		}
		dealer := model.NewDealer()
		player := model.NewPerfectBasicPlayer()
		fmt.Println(card, model.SimBlackjack(ds, dealer, player, SIM_ROUND))
	}
}

func main() {
	start := time.Now()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Dealer", "12", "13", "14", "15", "16", "17", "18", "19", "20", "Stand")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	fmt.Printf("Simulating Hard ")
	for du := 1; du <= 10; du++ {
		exp := []interface{}{du}
		var maxExp float64 = -100
		var bestStand int
		for stand := 12; stand < 21; stand++ {
			res := simluateBasicHard(du, stand)
			if res > maxExp {
				maxExp = res
				bestStand = stand
			}
			exp = append(exp, res)
		}
		exp = append(exp, bestStand)
		tbl.AddRow(exp...)
		fmt.Printf(".")
	}
	fmt.Println()
	tbl.Print()

	fmt.Println()
	fmt.Printf("Simulating Soft ")
	sTbl := table.New("Dealer", "15", "16", "17", "18", "19", "20", "Stand")
	sTbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for du := 1; du <= 10; du++ {
		exp := []interface{}{du}
		var maxExp float64 = -100
		var bestStand int
		for stand := 15; stand < 21; stand++ {
			res := simluateBasicSoft(du, stand)
			if res > maxExp {
				maxExp = res
				bestStand = stand
			}
			exp = append(exp, res)
		}
		exp = append(exp, bestStand)
		sTbl.AddRow(exp...)
		fmt.Printf(".")
	}
	fmt.Println()
	sTbl.Print()

	// simluateBasicHard(8, 17)

	// fmt.Println("Simulating Hard&Soft")
	// for du := 10; du <= 10; du++ {
	// 	var maxExp float64 = -100
	// 	var bestHard int
	// 	var bestSoft int
	// 	for hs := 12; hs < 21; hs++ {
	// 		for ss := 12; ss < 21; ss++ {
	// 			res := simulateBasicV2(du, hs, ss)
	// 			if res > maxExp {
	// 				maxExp = res
	// 				bestHard = hs
	// 				bestSoft = ss
	// 			}
	// 		}
	// 	}

	// 	fmt.Println(du, bestHard, bestSoft)
	// }
	// fmt.Println()

	// fmt.Println("Perfect basic player", simulatePerfectBasicPlayer())

	// simulatePerfectBasicPlayerBiasedShoe()

	elapsed := time.Since(start)
	fmt.Printf("Finished in %s\n", elapsed)
}
