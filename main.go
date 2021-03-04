package main

import (
	"fmt"
	"strings"

	"card-project/game"
)

func main() {
	//
	// newCards := deck.CreateDeck(1)
	// newCards.ShuffleDeck()
	// newCards.WriteToFile("./tmp/savedDeck")
	// loadedCards := deck.ReadDeckFromFile("./tmp/savedDeck")
	// if !reflect.DeepEqual(newCards, loadedCards) {
	// 	fmt.Println(reflect.DeepEqual(newCards, loadedCards))
	// 	panic("Something wrong with the decks")
	// }

	state := game.StartGame()

	game.MainLoop(&state)
	if state.Winner == "" {
		fmt.Println("Game Over!  It's a Tie!")
	}
	fmt.Println("Game Over!  Winners is:", strings.ToUpper(state.Winner))

}
