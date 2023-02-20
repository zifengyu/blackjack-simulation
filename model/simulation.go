package model

func processNature(dealer Gamer, player Gamer) (bool, float64) {
	if dealer.Value() != 21 && player.Value() != 21 {
		return false, 0
	}

	if dealer.Value() == 21 && player.Value() != 21 {
		return true, -1
	}

	return true, 1.5
}

func processResult(dealer Gamer, player Gamer) float64 {
	if dealer.IsBust() || player.Value() > dealer.Value() {
		return 1
	}

	if player.Value() < dealer.Value() {
		return -1
	}

	return 0
}

func SimBlackjack(ds *DealingShoe, dealer Gamer, player Gamer, round int) float64 {
	var totalMoney float64 = 0

	ds.Shuffle()
	for r := 0; r < round; r++ {
		if ds.RemainCards() < 10 {
			ds.Shuffle()
		}

		dealer.CleanCards()
		player.CleanCards()

		dealerUp := ds.Draw()

		dealer.GetCard(dealerUp)
		dealer.GetCard(ds.Draw())

		player.DealerUpCard(dealerUp)
		player.GetCard(ds.Draw())
		player.GetCard(ds.Draw())

		if isNature, m := processNature(dealer, player); isNature {
			totalMoney += m
			continue
		}

		// handle split card
		if player.CanSplit() && player.IsSplit(dealerUp) {
			for !dealer.IsStand() {
				dealer.GetCard(ds.Draw())
			}

			sc := player.Cards()[0]
			for i := 0; i < 2; i++ {
				player.CleanCards()
				player.GetCard(sc)
				player.GetCard(ds.Draw())

				for !player.IsStand() {
					player.GetCard(ds.Draw())
				}

				if player.IsBust() {
					totalMoney -= 1
					continue
				}

				totalMoney += processResult(dealer, player)
			}
		}

		var rate float64 = 1
		if player.IsDouble() {
			rate = 2
			player.GetCard(ds.Draw())
		} else {
			for !player.IsStand() {
				player.GetCard(ds.Draw())
			}
		}

		if player.IsBust() {
			totalMoney -= 1 * rate
			continue
		}

		for !dealer.IsStand() {
			dealer.GetCard(ds.Draw())
		}

		totalMoney += processResult(dealer, player) * rate
	}

	return totalMoney / float64(round)
}

func SimBasicStand(ds *DealingShoe, dealer Gamer, player Gamer, dealerUp int, round int) float64 {
	var totalMoney float64 = 0
	ds.Remove(dealerUp)

	for r := 0; r < round; {
		ds.Shuffle()

		dealer.CleanCards()
		player.CleanCards()

		dealer.GetCard(dealerUp)
		dealer.GetCard(ds.Draw())

		player.DealerUpCard(dealerUp)
		player.GetCard(ds.Draw())
		player.GetCard(ds.Draw())

		if isNature, _ := processNature(dealer, player); isNature {
			continue
		}

		for !player.IsStand() {
			player.GetCard(ds.Draw())
		}

		if player.IsBust() {
			totalMoney -= 1
			r++
			continue
		}

		for !dealer.IsStand() {
			dealer.GetCard(ds.Draw())
		}

		m := processResult(dealer, player)
		totalMoney += m
		r++
	}

	return totalMoney / float64(round)
}

func SimBasicStandHard(ds *DealingShoe, dealer Gamer, player Gamer, dealerUp int, round int) float64 {
	var totalMoney float64 = 0
	ds.Remove(dealerUp)

	for r := 0; r < round; {
		ds.Shuffle()

		dealer.CleanCards()
		player.CleanCards()

		dealer.GetCard(dealerUp)
		dealer.GetCard(ds.Draw())

		player.DealerUpCard(dealerUp)
		player.GetCard(ds.Draw())
		player.GetCard(ds.Draw())

		if isNature, _ := processNature(dealer, player); isNature {
			continue
		}

		if player.IsSoft() {
			continue
		}

		for !player.IsStand() {
			player.GetCard(ds.Draw())
			if player.IsSoft() {
				continue
			}
		}

		if player.IsBust() {
			totalMoney -= 1
			r++
			continue
		}

		for !dealer.IsStand() {
			dealer.GetCard(ds.Draw())
		}

		m := processResult(dealer, player)
		totalMoney += m
		r++
	}

	return totalMoney / float64(round)
}

func SimBasicStandSoft(ds *DealingShoe, dealer Gamer, player Gamer, dealerUp int, round int) float64 {
	var totalMoney float64 = 0
	ds.Remove(dealerUp)
	ds.Remove(1)

	for r := 0; r < round; {
		ds.Shuffle()

		dealer.CleanCards()
		player.CleanCards()

		dealer.GetCard(dealerUp)
		dealer.GetCard(ds.Draw())

		player.DealerUpCard(dealerUp)
		player.GetCard(1)
		player.GetCard(ds.Draw())

		if isNature, _ := processNature(dealer, player); isNature {
			continue
		}

		for !player.IsStand() {
			player.GetCard(ds.Draw())
		}

		if player.IsBust() {
			totalMoney -= 1
			r++
			continue
		}

		for !dealer.IsStand() {
			dealer.GetCard(ds.Draw())
		}

		m := processResult(dealer, player)
		totalMoney += m
		r++
	}

	return totalMoney / float64(round)
}

func SimDealerValue(ds *DealingShoe, dealer Gamer, dealerUp int, round int) []float64 {
	count := make([]float64, 6)

	ds.Remove(dealerUp)
	for r := 0; r < round; {
		ds.Shuffle()
		dealer.CleanCards()
		dealer.GetCard(dealerUp)
		dealer.GetCard(ds.Draw())
		if dealer.Value() == 21 {
			continue
		}

		for !dealer.IsStand() {
			dealer.GetCard(ds.Draw())
		}

		if dealer.IsBust() {
			count[5]++
		} else {
			count[dealer.Value()-17]++
		}
		r++
	}

	for i := 0; i < len(count); i++ {
		count[i] = count[i] / float64(round)
	}

	return count
}
