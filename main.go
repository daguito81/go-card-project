package main

import (
	"fmt"
	"os"
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

	game.MainLoop(&state, os.Stdin)
	if state.Winner == "" {
		fmt.Println("Game Over!  It's a Tie!")
	}
	fmt.Println("Game Over!  Winners is:", strings.ToUpper(state.Winner))
	fmt.Println("##### Player ######")
	state.Player.Print()
	fmt.Printf("Total Points: %d\n", state.PlayerScore)
	fmt.Println("###### House ######")
	state.House.Print()
	fmt.Printf("Total Points: %d\n", state.HouseScore)

}
