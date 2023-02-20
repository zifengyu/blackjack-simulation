package model

type Gamer interface {
	CleanCards()
	Cards() []int
	GetCard(card int)
	DealerUpCard(card int)
	Value() int
	// HasAce() bool
	CanSplit() bool
	IsSoft() bool
	IsBust() bool
	IsDouble() bool
	IsSplit(card int) bool
	IsStand() bool
}

type handCards []int

func (hc *handCards) CleanCards() {
	*hc = (*hc)[:0]
}

func (hc *handCards) Cards() []int {
	return *hc
}

func (hc *handCards) GetCard(card int) {
	*hc = append(*hc, card)
}

func (hc *handCards) HasAce() bool {
	for _, card := range *hc {
		if card == 1 {
			return true
		}
	}
	return false
}

func (hc *handCards) Value() int {
	value := 0
	for _, card := range *hc {
		value += card
	}

	if hc.HasAce() && value <= 11 {
		value += 10
	}
	return value
}

func (hc handCards) CanSplit() bool {
	return len(hc) == 2 && hc[0] == hc[1]
}

func (hc *handCards) IsSoft() bool {
	if !hc.HasAce() {
		return false
	}

	value := 0
	for _, card := range *hc {
		value += card
	}

	return value <= 11
}

func (hc *handCards) IsBust() bool {
	return hc.Value() > 21
}

type Dealer struct {
	handCards
}

func NewDealer() Gamer {
	return &Dealer{}
}

func (d *Dealer) IsStand() bool {
	// if d.Value() > 17 {
	// 	return true
	// }

	// if d.Value() == 17 && d.IsSoft() {
	// 	return true
	// }

	// return false
	return d.Value() >= 17
}

func (d *Dealer) IsDouble() bool {
	return false
}

func (d *Dealer) IsSplit(card int) bool {
	return false
}

func (d *Dealer) DealerUpCard(card int) {

}

type neverBustPlayer struct {
	handCards
}

func NewNeverBustPlayer() Gamer {
	return &neverBustPlayer{}
}

func (p *neverBustPlayer) IsStand() bool {
	if p.HasAce() {
		return p.Value() >= 18
	}

	return p.Value() >= 12
}

func (p *neverBustPlayer) IsDouble() bool {
	return false
}

func (p *neverBustPlayer) IsSplit(card int) bool {
	return false
}

func (p *neverBustPlayer) DealerUpCard(card int) {}

type basicStandPlayer struct {
	handCards
	stand int
}

func NewBasicStandPlayer(stand int) Gamer {
	return &basicStandPlayer{stand: stand}
}

func (p *basicStandPlayer) IsStand() bool {
	return p.Value() >= p.stand
}

func (p *basicStandPlayer) IsDouble() bool {
	return false
}

func (p *basicStandPlayer) IsSplit(card int) bool {
	return false
}

func (p *basicStandPlayer) DealerUpCard(card int) {}

type basicSoftStandPlayer struct {
	handCards
	stand int
	dc    int
}

func NewBasicSoftStandPlayer(stand int) Gamer {
	return &basicSoftStandPlayer{stand: stand}
}

func (p *basicSoftStandPlayer) IsDouble() bool {
	return false
}

func (p *basicSoftStandPlayer) IsStand() bool {
	if p.IsSoft() {
		return p.Value() >= p.stand
	}

	var stand int
	switch p.dc {
	case 2, 3:
		stand = 13
	case 4, 5, 6:
		stand = 12
	default:
		stand = 17
	}

	return p.Value() >= stand
}

func (p *basicSoftStandPlayer) IsSplit(card int) bool {
	return false
}

func (p *basicSoftStandPlayer) DealerUpCard(card int) {
	p.dc = card
}

type basicStandPlayerV2 struct {
	handCards
	hardStand int
	softStand int
}

func NewBasicPlayer(hard int, soft int) Gamer {
	return &basicStandPlayerV2{hardStand: hard, softStand: soft}
}

func (p *basicStandPlayerV2) IsStand() bool {
	if p.IsSoft() {
		return p.Value() >= p.softStand
	}

	return p.Value() >= p.hardStand
}

func (p *basicStandPlayerV2) IsDouble() bool {
	return false
}

func (p *basicStandPlayerV2) IsSplit(card int) bool {
	return false
}

func (p *basicStandPlayerV2) DealerUpCard(card int) {
}

type perfectBasicPlayer struct {
	handCards
	dealerUp int
}

func NewPerfectBasicPlayer() Gamer {
	return &perfectBasicPlayer{}
}

func (p *perfectBasicPlayer) IsDouble() bool {
	if p.IsSoft() {
		c := p.Cards()[0]
		if c == 1 {
			c = p.Cards()[1]
		}

		switch c {
		case 7:
			return p.dealerUp == 3 || p.dealerUp == 4 || p.dealerUp == 5 || p.dealerUp == 6
		case 6:
			return p.dealerUp >= 2 && p.dealerUp <= 6
		case 2, 3, 4, 5:
			return p.dealerUp == 4 || p.dealerUp == 5 || p.dealerUp == 6
		case 1:
			return p.dealerUp == 5 || p.dealerUp == 6
		}

		return false
	}

	switch p.Value() {
	case 11:
		return true
	case 10:
		return p.dealerUp != 1 && p.dealerUp != 10
	case 9:
		return p.dealerUp != 1 && p.dealerUp <= 6
	case 8:
		return (p.dealerUp == 5 || p.dealerUp == 6) && p.Cards()[0] != 6 && p.Cards()[0] != 2
	}

	return false
}

func (p *perfectBasicPlayer) IsStand() bool {
	if p.IsSoft() {
		// soft stand
		switch p.dealerUp {
		case 1, 2, 3, 4, 5, 6, 7, 8:
			return p.Value() >= 18
		case 9, 10:
			return p.Value() >= 19
		}
	}

	// hard stand
	switch p.dealerUp {
	case 2, 3:
		return p.Value() >= 13
	case 4, 5, 6:
		return p.Value() >= 12
	case 1, 7, 8, 9, 10:
		return p.Value() >= 17
	}

	panic("should not reach here")
}

func (p *perfectBasicPlayer) IsSplit(card int) bool {
	switch p.Cards()[0] {
	case 1, 8:
		return true
	case 2, 3, 6:
		return card >= 2 && card <= 7
	case 4:
		return card == 5
	case 5, 10:
		return false
	case 7:
		return card >= 2 && card <= 8
	case 9:
		return (card >= 2 && card <= 6) || card == 8 || card == 9
	}
	return false
}

func (p *perfectBasicPlayer) DealerUpCard(card int) {
	p.dealerUp = card
}
