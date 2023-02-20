package main

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/zifengyu/blackjack-simulation/model"
)

type State struct {
	value int
	soft  bool

	upStates   []chan<- float64
	downStates []<-chan float64
	downProb   []float64
	stand      bool
	ev         float64
	sEv        float64
	hEv        float64
}

func (st *State) Cal() {
	for i, c := range st.downStates {
		dev := <-c
		st.hEv += st.downProb[i] * dev
	}

	if len(st.downStates) > 0 && st.hEv > st.sEv {
		st.ev = st.hEv
		st.stand = false
	} else {
		st.ev = st.sEv
		st.stand = true
	}

	for _, c := range st.upStates {
		c <- st.ev
		close(c)
	}
}

func (st *State) Print() {
	if st.soft {
		fmt.Print("S")
	} else {
		fmt.Print("H")
	}
	fmt.Printf("%d: sEv=%f hEv=%f\n", st.value, st.sEv, st.hEv)
}

func makeState(value int, soft bool, dealerResult []float64) *State {
	// stand expected value
	var sEv float64 = dealerResult[5]
	for dealer := 17; dealer <= 21; dealer++ {
		if dealer > value {
			sEv -= dealerResult[dealer-17]
		} else if dealer < value {
			sEv += dealerResult[dealer-17]
		}
	}
	return &State{
		value: value,
		soft:  soft,
		stand: true,
		sEv:   sEv,
	}
}

func createEdge(from *State, to *State, prob float64) {
	edge := make(chan float64, 1)
	from.downStates = append(from.downStates, edge)
	from.downProb = append(from.downProb, prob)
	to.upStates = append(to.upStates, edge)
}

func createGraph(decks float64, upCard int, dealerResult []float64) ([]*State, []*State) {
	totalCards := 52 * decks
	hardStates := make([]*State, 22)
	softStates := make([]*State, 22)
	for i := 2; i <= 21; i++ {
		hardStates[i] = makeState(i, false, dealerResult)
		softStates[i] = makeState(i, true, dealerResult)
	}

	bustState := &State{
		value: 22,
		sEv:   -1,
		hEv:   -1,
	}

	for i := 2; i <= 21; i++ {
		for card := 2; card <= 10; card++ {
			var cardProb float64 = 0
			if card != 10 {
				cardProb = 4.0 * decks / float64(totalCards-1)
			} else {
				cardProb = 16.0 * decks / float64(totalCards-1)
			}

			if card == upCard {
				cardProb -= 1.0 / float64(totalCards-1)
			}
			if i+card <= 21 {
				createEdge(hardStates[i], hardStates[i+card], cardProb)
				createEdge(softStates[i], softStates[i+card], cardProb)
			} else {
				createEdge(hardStates[i], bustState, cardProb)
				createEdge(softStates[i], hardStates[i+card-10], cardProb)
			}
		}

		var aceProb float64 = 4.0 * decks / float64(totalCards-1)
		if upCard == 1 {
			aceProb -= 1.0 / float64(totalCards-1)
		}
		if i+11 <= 21 {
			createEdge(hardStates[i], softStates[i+11], aceProb)
			createEdge(softStates[i], softStates[i+11], aceProb)
		} else if i+1 <= 21 {
			createEdge(hardStates[i], hardStates[i+1], aceProb)
			createEdge(softStates[i], softStates[i+1], aceProb)
		} else {
			createEdge(hardStates[i], bustState, aceProb)
			createEdge(softStates[i], hardStates[i-9], aceProb)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(40)
	for i := 2; i <= 21; i++ {
		go func(v int) {
			hardStates[v].Cal()
			wg.Done()
		}(i)

		go func(v int) {
			softStates[v].Cal()
			wg.Done()
		}(i)
	}

	bustState.Cal()
	wg.Wait()

	return hardStates, softStates
}

func calDealerValue(round int, print bool) [][]float64 {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Dealer", "17", "18", "19", "20", "21", "Busting")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	result := make([][]float64, 10)
	for d := 1; d <= 10; d++ {
		ds := model.NewDealingShoe(1)
		dealer := model.NewDealer()
		result[d-1] = model.SimDealerValue(ds, dealer, d, round)
		tbl.AddRow(d, result[d-1][0], result[d-1][1], result[d-1][2], result[d-1][3], result[d-1][4], result[d-1][5])
	}

	if print {
		fmt.Println("Dealer Card Probability")
		tbl.Print()
	}
	return result
}

func main() {
	dealerResult := calDealerValue(10000000, true)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	hTbl := table.New("Dealer", "12", "13", "14", "15", "16", "17", "18", "19", "20")
	hTbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	sTbl := table.New("Dealer", "12", "13", "14", "15", "16", "17", "18", "19", "20")
	sTbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for dealerUp := 2; dealerUp <= 11; dealerUp++ {
		du := dealerUp
		if du == 11 {
			du = 1
		}
		hardStates, softStates := createGraph(6, du, dealerResult[du-1])

		hard := []interface{}{dealerUp}
		soft := []interface{}{dealerUp}
		for i := 12; i <= 20; i++ {
			if hardStates[i].stand {
				hard = append(hard, "S")
			} else {
				hard = append(hard, "H")
			}

			if softStates[i].stand {
				soft = append(soft, "S")
			} else {
				soft = append(soft, "H")
			}
		}

		hTbl.AddRow(hard...)
		sTbl.AddRow(soft...)
	}

	fmt.Println("Hard Strategy")
	hTbl.Print()

	fmt.Println("Soft Strategy")
	sTbl.Print()

	_, softStates := createGraph(52, 1, dealerResult[0])
	softStates[17].Print()
	softStates[18].Print()
	softStates[19].Print()
}
