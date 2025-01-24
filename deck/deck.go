package deck

import "GoJack/card"

type Deck struct {
	cards card.Card[52]
}
