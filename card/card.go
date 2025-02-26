package card

import "fmt"

const HeartString = "\u2665"
const DiamondString = "\u2666"
const ClubString = "\u2663"
const SpadesString = "\u2660"

const (
	HEART   uint8 = 0
	CLUB    uint8 = 1
	DIAMOND uint8 = 2
	SPADE   uint8 = 3
)

type Card struct {
	Num  uint8
	Suit uint8
}

func (c Card) ToStr() string {
	var s string
	var suitStr string

	switch c.Suit {
	case HEART:

		suitStr = HeartString
	case DIAMOND:
		suitStr = DiamondString
	case CLUB:

		suitStr = ClubString
	case SPADE:
		suitStr = SpadesString
	default:
		suitStr = "ERROR"
	}
	switch c.Num {
	case 11:
		s = fmt.Sprintf("J")
		break
	case 12:
		s = fmt.Sprintf("Q")
		break
	case 13:
		s = fmt.Sprintf("K")
		break
	default:
		s = fmt.Sprintf("%d", c.Num)
		break
	}
	s += fmt.Sprintf(" %s", suitStr)
	return s
}
