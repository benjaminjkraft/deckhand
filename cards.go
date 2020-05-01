package cards

// basic data structures, nothing here is specific to my solution

import (
	"fmt"
	"math/rand"
)

type Suit uint8

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
	NumSuits Card = iota
)

type Rank uint8

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	NumRanks Card = iota
)

const DeckSize = NumSuits * NumRanks

type Card uint8

func (c Card) Suit() Suit { return Suit(c / NumRanks) }
func (c Card) Rank() Rank { return Rank(c % NumRanks) }
func (c Card) String() string {
	if c >= Card(DeckSize) {
		return fmt.Sprintf("?%d", c)
	}

	// https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
	rankForUnicode := c.Rank()
	if rankForUnicode >= Queen {
		// unicode has Knight between Jack and Queen
		rankForUnicode += 1
	}

	return string(rune(0x1f0a1) + rune(rankForUnicode) + rune(0x10*c.Suit()))
}

type Deck [DeckSize]Card

func NewDeck() Deck {
	var d Deck
	for i := range d {
		d[i] = Card(i)
	}
	return d
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}
