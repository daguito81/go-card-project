package game

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"card-project/deck"
)

type State struct {
	Active       bool
	Main         deck.Deck
	Player       deck.Deck
	PlayerScore  int
	PlayerActive bool
	PlayerBust   bool
	House        deck.Deck
	HouseScore   int
	HouseBust    bool
	Winner       string
}

func StartGame() State {
	fmt.Println("Starting New Game!")
	mainDeck := deck.CreateDeck(1)
	mainDeck.ShuffleDeck()
	gs := State{
		Active:       true,
		Main:         mainDeck,
		Player:       make(deck.Deck, 0, 640),
		House:        make(deck.Deck, 0, 640),
		PlayerScore:  0,
		HouseScore:   0,
		PlayerActive: true,
	}

	gs.Player = append(gs.Player, gs.Main.DealCards(2)...)
	gs.PlayerScore = gs.Player.GetValue()
	gs.House = append(gs.House, gs.Main.DealCards(2)...)
	gs.HouseScore = gs.House.GetValue()

	return gs
}

func MainLoop(gs *State, r io.Reader) {
	for gs.Active {
		playerLoop(gs, r)
		if !gs.Active {
			break
		}
		houseLoop(gs)
		gs.printState(gs.PlayerActive)
	}
	gs.finalCheck()

}

func playerLoop(gs *State, r io.Reader) {
	for gs.PlayerActive {
		gs.printState(gs.PlayerActive)
		fmt.Println("Choose Hit, Stand or Exit")
		choice, err := getChoice(r)
		if err != nil {
			fmt.Println("Something went wrong, choose again")
			continue
		}
		switch choice {
		case "hit":
			gs.Player = append(gs.Player, gs.Main.DealCards(1)...)
			gs.update()
		case "stand":
			gs.rollingCheck()
			gs.PlayerActive = false
		case "exit":
			gs.PlayerActive = false
			os.Exit(0)
		}
	}
}

func houseLoop(gs *State) {
	for gs.HouseScore <= 16 {
		gs.House = append(gs.House, gs.Main.DealCards(1)...)
		gs.printState(gs.PlayerActive)
		gs.update()
		gs.rollingCheck()
		time.Sleep(3 * time.Second)
	}
	gs.Active = false
}

func getChoice(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	fmt.Print("->")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error getting Choice from Stdin: %s", err)
		return "ERROR", err
	}
	text = strings.Replace(text, "\n", "", -1) // Need \r\n for windows?
	text = strings.Replace(text, "\r", "", -1) // Need \r\n for windows?

	return strings.ToLower(text), nil
}

func (gs *State) printState(hiddenHouse bool) {
	fmt.Println("Game State:")
	fmt.Println("------ Player Hand ------")
	gs.Player.Print()
	fmt.Println("------ House Hand------")
	if hiddenHouse {
		fmt.Println([]string{gs.House[0]})
	} else {
		fmt.Println(gs.House)
	}
}

func (gs *State) update() {
	gs.PlayerScore = gs.Player.GetValue()
	gs.HouseScore = gs.House.GetValue()
	gs.rollingCheck()

}

func (gs *State) rollingCheck() {
	if gs.PlayerScore > 21 {
		fmt.Println("Player bust!", gs.PlayerScore)
		gs.PlayerBust = true
		gs.PlayerActive = false
	}
	if gs.HouseScore > 21 {
		fmt.Println("House bust!", gs.HouseScore)
		gs.HouseBust = true
	}
	if gs.PlayerBust && !gs.HouseBust {
		gs.Winner = "house"
	} else if gs.HouseBust && !gs.PlayerBust {
		gs.Winner = "player"
	} else {
		gs.Winner = ""
	}

	if gs.Winner != "" {
		gs.Active = false
	}

}

func (gs *State) finalCheck() {
	if gs.PlayerBust && gs.HouseBust {
		gs.Winner = ""
		gs.Active = false
		return
	}
	if gs.PlayerBust || gs.HouseBust {
		return
	}
	if gs.PlayerScore > gs.HouseScore {
		gs.Winner = "player"
	} else if gs.HouseScore > gs.PlayerScore {
		gs.Winner = "house"
	} else {
		gs.Winner = ""
	}

	gs.Active = false
}
