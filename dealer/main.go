package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/zifengyu/blackjack-simulation/model"
)

const SIM_ROUND = 1000000

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Dealer", "17", "18", "19", "20", "21", "Busting")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	result := make([][]float64, 10)
	for d := 1; d <= 10; d++ {
		ds := model.NewDealingShoe(1)
		dealer := model.NewDealer()
		result[d-1] = model.SimDealerValue(ds, dealer, d, SIM_ROUND)
		tbl.AddRow(d, result[d-1][0], result[d-1][1], result[d-1][2], result[d-1][3], result[d-1][4], result[d-1][5])
	}
	tbl.Print()
}
