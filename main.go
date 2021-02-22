package main

import (
	"fmt"
	"time"

	"card-project/deck"
)

func main() {

	newCards := deck.CreateDeck()
	// newCards.Print()
	newCards.ShuffleDeck()
	hand1 := newCards.DealCards(2)
	hand2 := newCards.DealCards(2)
	hand3 := newCards.DealCards(2)

	fmt.Println("Hand 1")
	hand1.Print()
	fmt.Println("Hand 2")
	hand2.Print()
	fmt.Println("Hand 3")
	hand3.Print()
	time.Sleep(5*time.Second)

	table := newCards.DealCards(3) // Flop
	fmt.Println("FLOP")
	table.Print()
	time.Sleep(time.Second)
	table = append(table, newCards.DealCards(1)[0]) // Turn
	fmt.Println("TURN")
	table.Print()
	time.Sleep(time.Second)

	table = append(table, newCards.DealCards(1)[0]) // River
	fmt.Println("RIVER")
	table.Print()
	time.Sleep(time.Second)


}
