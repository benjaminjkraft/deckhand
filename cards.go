package cards

import (
	"math/big"
	"math/rand"
)

type Suit uint8

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
	NumSuits uint8 = iota
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
	NumRanks uint8 = iota
)

const DeckSize = NumSuits * NumRanks

type Card uint8

func (c Card) Suit() Suit { return Suit(uint8(c) / NumRanks) }
func (c Card) Rank() Rank { return Rank(uint8(c) % NumRanks) }
func (c Card) String() string {
	// https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
	return string(rune(0x1f0a1) + rune(c.Rank()) + 0x10*rune(c.Suit()))
}

type Deck [DeckSize]Card

func NewDeck() Deck {
	var d Deck
	for i := range d {
		d[i] = Card(i)
	}
	return d
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func (d Deck) Index() *big.Int {
	retval := big.NewInt(0)
	for i, c := range d {
		// retval *= i+1
		retval.Mul(retval, big.NewInt(int64(i+1)))
		// retval += c
		retval.Add(retval, big.NewInt(int64(c)))
	}
	return retval
}

func FromIndex(index *big.Int) Deck {
	var d Deck
	for ri := range d {
		i := len(d) - ri

		// index, c = index / i, index % i
		c := big.NewInt(0)
		index, c = index.DivMod(index, big.NewInt(int64(i)), c)

		d[i-1] = Card(c.Int64())
	}
	return d
}

func IndexEncode(d Deck) []byte {
	return d.Index().Bytes()
}

func IndexDecode(b []byte) Deck {
	return FromIndex(big.NewInt(0).SetBytes(b))
}
