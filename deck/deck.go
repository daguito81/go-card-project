package deck

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'Deck'
// which is a slice of strings
type Deck []string

// Setting up a style of ENUM
func getCardSuits() [4]string {
	return [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}
}
func getCardValues() [13]string {
	return [13]string{"Ace", "Two", "Three", "Four",
		"Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King"}
}

// CreateDeck will return a Deck value with 52 cards based on a standard deck of cards
func CreateDeck(numDecks int) Deck {
	var newDeck Deck
	for i := 0; i < numDecks; i++ {
		for _, suit := range getCardSuits() {
			for _, value := range getCardValues() {
				newDeck = append(newDeck, fmt.Sprintf("%s of %s", value, suit))
			}
		}
	}
	fmt.Println("Created Deck with", numDecks, "packs of cards")

	return newDeck
}

// Print will print the contents of the deck in index/value format to stdout
func (d Deck) Print() {
	for i, card := range d {
		fmt.Printf("%d %s\n", i+1, card)
	}
}

// DealCards will deal the N amount of cards from the top of the deck
func (d *Deck) DealCards(cardsPerHand int) Deck {
	if cardsPerHand > len(*d) {
		log.Fatalln("Not enough cards to deal")
	}
	newHand := (*d)[:cardsPerHand]
	*d = (*d)[cardsPerHand:]
	return newHand
}

// Shuffle will randomly change the order of the cards in the Deck
func (d *Deck) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Shuffling")
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})

}

func (d *Deck) WriteToFile(path string) {
	data := *d
	fi, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	var dataBytes [][]byte
	for _, val := range data {
		dataBytes = append(dataBytes, []byte(val))
	}

	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(dataBytes); err != nil {
		panic(err)
	}

	if err := os.WriteFile(fi.Name(), buf.Bytes(), 0644); err != nil {
		panic(err)
	}

}

func ReadDeckFromFile(path string) Deck {
	var cardsBytes [][]byte
	var cards Deck
	fi, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	dec := gob.NewDecoder(fi)

	if err := dec.Decode(&cardsBytes); err != nil {
		panic(err)
	}

	for _, card := range cardsBytes {
		cards = append(cards, string(card))
	}

	return cards
}

func (d *Deck) GetValue() int {
	var values []int
	var aceCount int

	for _, val := range *d {
		switch strings.Fields(val)[0] {
		case "Ace":
			aceCount += 1
			values = append(values, 11)
		case "Two":
			values = append(values, 2)
		case "Three":
			values = append(values, 3)
		case "Four":
			values = append(values, 4)
		case "Five":
			values = append(values, 5)
		case "Six":
			values = append(values, 6)
		case "Seven":
			values = append(values, 7)
		case "Eight":
			values = append(values, 8)
		case "Nine":
			values = append(values, 9)
		case "Ten":
			values = append(values, 10)
		case "Jack":
			values = append(values, 10)
		case "Queen":
			values = append(values, 10)
		case "King":
			values = append(values, 10)
		}
		// "Two", "Three", "Four",
		// 	"Five", "Six", "Seven", "Eight",
		// 	"Nine", "Ten", "Jack", "Queen", "King"
	}
	total := 0
	for _, val := range values {
		total += val
	}
	for aceCount > 0 && total > 21 {
		total -= 10
		aceCount -= 1
	}
	return total
}
