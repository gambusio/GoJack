package deck

import "GoJack/card"

type Deck struct {
	cards [52]card.Card
}

func NewDeck() Deck {
	var d = new(Deck)
}

func (Deck Deck) Shuffle() {
	//shuffle the deck
}
