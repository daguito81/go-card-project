package main

import (
	"fmt"
	"os"
	"strings"

	"card-project/game"
)

func main() {

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
