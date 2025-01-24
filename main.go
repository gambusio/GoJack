package main

import (
	"GoJack/card"
	"fmt"
)

func main() {
	carta1 := card.Card{1, card.SPADE}
	fmt.Println(carta1.ToStr())
}
