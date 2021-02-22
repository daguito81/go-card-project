package deck

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const casinoDecks = 6
const cardsPerDeck = 52

// Create Deck

func Test_Deck_CreateDeck_Single(t *testing.T) {

	got := CreateDeck(1)
	// Decks are always ordered the same when when created
	if got[0] != "Ace of Hearts" {
		t.Errorf("First Card should be Ace of Hearts, got: %s", got[0])
	}
	if got[30] != "Five of Clubs" {
		t.Errorf("First Card should be Ace of Hearts, got: %s", got[30])
	}
	if got[15] != "Three of Diamonds" {
		t.Errorf("Card should be Three of Diamonds, got: %s", got[15])
	}
	if got[len(got)-1] != "King of Spades" {
		t.Errorf("Last Card should be King of Spades, got: %s", got[len(got)-1])
	}
	if len(got) != cardsPerDeck {
		t.Errorf("New Deck should have 52 cards, got: %d", len(got))
	}

}
func Test_Deck_CreateDeck_Multi(t *testing.T) {
	newDeck := CreateDeck(casinoDecks)
	if len(newDeck) != cardsPerDeck*casinoDecks {
		t.Errorf("New Deck should have %d cards, got: %d", casinoDecks*cardsPerDeck, len(newDeck))
	}
	if newDeck[0] != newDeck[52] {
		t.Errorf("Decks are not in the right order: %s -> %s", newDeck[0], newDeck[52])
	}
}

func Benchmark_Deck_CreateDeck(b *testing.B) {
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < b.N; i++ {
		CreateDeck(1)
	}
	os.Stdout = rescueStdout
}

// Print

func Test_Deck_Print(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	deck := Deck{"Ace of Spades", "Queen of Hearts"}
	deck.Print()

	if err := w.Close(); err != nil {
		panic(err)
	}

	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	want := fmt.Sprint("1 Ace of Spades\n2 Queen of Hearts\n")
	if string(out) != want {
		t.Errorf("Test Failed. Wanted: %s Got: %s", want, string(out))
	}
}

func Benchmark_Deck_Print(b *testing.B) {
	deck := CreateDeck(1)
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < b.N; i++ {
		deck.Print()
	}
	os.Stdout = rescueStdout
}

// Deal Cards
func Test_Deck_DealCards(t *testing.T) {
	const cardsPerHand = 5
	deck := CreateDeck(1)
	originalDeckSize := len(deck)

	hand1 := deck.DealCards(cardsPerHand)
	if len(hand1) != cardsPerHand || len(deck) != originalDeckSize-cardsPerHand {
		t.Errorf("Wrong Hand/Deck Size (wanted 5) - Hand: %d Deck: %d", len(hand1), len(deck))
	}
	firstCardHand1 := hand1[0]
	if firstCardHand1 != "Ace of Hearts" {
		t.Errorf("Wrong Card Dealt. Wanted: Ace of Hearts Got:%s", firstCardHand1)
	}

	hand2 := deck.DealCards(cardsPerHand)
	if len(hand2) != cardsPerHand || len(deck) != originalDeckSize-cardsPerHand*2 {
		t.Errorf("Wrong Hand/Deck Size (wanted 5) - Hand: %d Deck: %d", len(hand2), len(deck))
	}
	firstCardHand2 := hand2[0]
	if firstCardHand2 != "Six of Hearts" {
		t.Errorf("Wrong Card Dealt. Wanted: Six of Hearts Got:%s", firstCardHand2)
	}

	hand3 := deck.DealCards(cardsPerHand)
	if len(hand3) != cardsPerHand || len(deck) != originalDeckSize-cardsPerHand*3 {
		t.Errorf("Wrong Hand/Deck Size (wanted 5) - Hand: %d Deck: %d", len(hand3), len(deck))
	}
	firstCardHand3 := hand3[0]
	if firstCardHand3 != "Jack of Hearts" {
		t.Errorf("Wrong Card Dealt. Wanted: Jack of Hearts Got:%s", firstCardHand3)
	}
}

func Benchmark_Deck_DealCards(b *testing.B) {
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < b.N; i++ {
		deck := CreateDeck(1)
		deck.DealCards(5)
	}
	os.Stdout = rescueStdout

}

func Test_Deck_ShuffleDeck(t *testing.T) {
	cards := CreateDeck(1)
	backupCards := CreateDeck(1)
	cards.ShuffleDeck()
	if reflect.DeepEqual(cards, backupCards) {
		t.Errorf("Slices are equal, meaning there was no shuffling: %s -> %s", backupCards[1], cards[1])
	}
}

func Benchmark_Deck_ShuffleDeck(b *testing.B) {
	cards := CreateDeck(1)
	rescueStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < b.N; i++ {
		cards.ShuffleDeck()
	}
	os.Stdout = rescueStdout
}
