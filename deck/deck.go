package deck

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"os"
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

	return newDeck
}

// Print will print the contents of the deck in index/value format to stdout
func (d Deck) Print() {
	for i, card := range d {
		fmt.Printf("%d %s\n", i+1, card)
	}
}

// DealCards will deal the N ammount of cards from the top of the deck
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
