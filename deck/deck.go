package deck

import (
	"fmt"
	"log"
	"math/rand"
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
