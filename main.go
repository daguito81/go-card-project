package main

import (
	"fmt"
	"reflect"
	"time"

	"card-project/deck"
)

func main() {

	newCards := deck.CreateDeck(1)
	newCards.ShuffleDeck()
	newCards.WriteToFile("./tmp/savedDeck")
	loadedCards := deck.ReadDeckFromFile("./tmp/savedDeck")
	if !reflect.DeepEqual(newCards, loadedCards) {
		fmt.Println(reflect.DeepEqual(newCards, loadedCards))
		panic("Something wrong with the decks")
	}

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
	time.Sleep(5 * time.Second)

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
