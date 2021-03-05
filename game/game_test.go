package game

import (
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_Game_StartGame(t *testing.T) {

	got := StartGame()

	if got.Active != true {
		t.Errorf("Game is starting with Active = False")
	}

	if len(got.Player) != 2 {
		t.Errorf("Player Hand should have 2 cards, has %d instead", len(got.Player))
	}

	if len(got.House) != 2 {
		t.Errorf("House Hand should have 2 cards, has %d instead", len(got.House))
	}

	if got.PlayerScore < 2 || got.HouseScore < 2 {
		t.Errorf("Something is wrong with hand value calculation: %d, %d", got.PlayerScore, got.HouseScore)
	}

	if got.PlayerBust || got.HouseBust {
		t.Errorf("Player and House should not be 'bust' at the start")
	}

}

func Benchmark_Game_StartGame(b *testing.B) {
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < b.N; i++ {
		StartGame()
	}
	os.Stdout = rescueStdout
}

func Test_Game_getChoice(t *testing.T) {
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)

	want := "Hit"

	got, err := getChoice(strings.NewReader(want + "\n"))
	if err != nil {
		t.Errorf("Error running function")
	}
	if got != strings.ToLower(want) {
		t.Errorf("Got: %s Want: %s", got, strings.ToLower(want))
	}

	want = "StAnD"

	got, err = getChoice(strings.NewReader(want + "\n"))
	if err != nil {
		t.Errorf("Error running function")
	}
	if got != strings.ToLower(want) {
		t.Errorf("Got: %s Want: %s", got, strings.ToLower(want))
	}

	want = "exit"

	got, err = getChoice(strings.NewReader(want + "\n"))
	if err != nil {
		t.Errorf("Error running function")
	}
	if got != strings.ToLower(want) {
		t.Errorf("Got: %s Want: %s", got, strings.ToLower(want))
	}

	// Error
	want = "DaGo"

	got, err = getChoice(strings.NewReader(want))
	if err == nil {
		t.Errorf("Should've caught the error")
	}

	os.Stdout = rescueStdout

}

func Benchmark_Game_getChoice(b *testing.B) {
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testString := strings.NewReader("TestingChoiceString\n")
		_, err := getChoice(testString)
		if err != nil {
			b.Errorf("There was an error: %s", err)
		}
	}
	os.Stdout = rescueStdout
}

func Test_Game_finalCheck(t *testing.T) {
	state := StartGame()
	// Situation a player busts
	state.PlayerBust = true
	state.HouseBust = false
	state.finalCheck()
	if !state.Active {
		t.Errorf("Something ended the game too early")
	}
	if state.Winner != "" {
		t.Errorf("Something declared a winner prematurely")
	}

	// Situation 2 Player Wins

	state = StartGame()

	state.PlayerScore = 20
	state.HouseScore = 19

	state.finalCheck()
	if state.Winner != "player" {
		t.Errorf("Player should've won but winner was %s", state.Winner)
	}

	// Situation 3 House Wins
	state.HouseScore = 20
	state.PlayerScore = 19

	state.finalCheck()
	if state.Winner != "house" {
		t.Errorf("House should've won but winner was %s", state.Winner)
	}

	// Situation 3 House Wins
	state.HouseScore = 20
	state.PlayerScore = 20
	state.Winner = ""

	state.finalCheck()
	if state.Winner != "" {
		t.Errorf("Should've been a tie but someone won %s", state.Winner)
	}

	// Situation 4  Both Bust

	state = StartGame()
	state.PlayerBust = true
	state.HouseBust = true
	state.finalCheck()
	if state.Winner != "" {
		t.Errorf("Should've been a tie but someone won %s", state.Winner)
	}
}

func Benchmark_Game_finalCheck(b *testing.B) {
	state := StartGame()
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		state.PlayerScore = rand.Intn(30)
		if state.PlayerScore > 21 {
			state.PlayerBust = true
		}
		state.HouseScore = rand.Intn(30)
		if state.HouseScore > 21 {
			state.HouseBust = true
		}
		state.finalCheck()
	}
}

func Test_Game_rollingCheck(t *testing.T) {
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)

	// Situation 1 Player Bust
	state := StartGame()

	state.PlayerScore = 23
	state.rollingCheck()

	if state.PlayerBust == false || state.PlayerActive == true {
		t.Error("Player should've busted", state)
	}
	if state.Winner != "house" {
		t.Error("House should've won because player busted")
	}

	// Situation 2 House Busts
	state = StartGame()

	state.HouseScore = 22
	state.rollingCheck()

	if state.HouseBust == false {
		t.Error("House should've busted")
	}
	if state.Winner != "player" {
		t.Error("Player should've won because house busted")
	}

	// Situation 3 Tie
	state = StartGame()

	state.PlayerScore = 19
	state.HouseScore = 19

	state.rollingCheck()

	if state.HouseBust || state.PlayerBust {
		t.Error("Neither should've busted but they did", state)
	}
	if state.Winner != "" {
		t.Error("Winner declared prematurely")
	}
	os.Stdout = rescueStdout
}

func Benchmark_Game_rollingCheck(b *testing.B) {
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)

	state := StartGame()
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		state.PlayerScore = rand.Intn(30)
		if state.PlayerScore > 21 {
			state.PlayerBust = true
		}
		state.HouseScore = rand.Intn(30)
		if state.HouseScore > 21 {
			state.HouseBust = true
		}
		state.rollingCheck()
	}
	os.Stdout = rescueStdout
}
