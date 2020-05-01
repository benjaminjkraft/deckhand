package cards

import (
	"math/big"
)

// cardIndices computes the index of each card in the set of cards that are not
// before it in the deck.  So the first card's index is just its index in the
// deck; the last card's index is always 0.
//
// Note that the return value satisfies 0 <= indices[i] < 52-i for all i.
func (d Deck) cardIndices() [DeckSize]Card {
	var indices [DeckSize]Card
	var used [DeckSize]bool
	for i, c := range d {
		var r Card
		for _, u := range used[:c] {
			if !u {
				r++
			}
		}
		used[c] = true
		indices[i] = r
	}
	return indices
}

// Index computes the index of this permutation of the deck in the
// lexicographic ordering of all permutations of the deck.
//
// Specifically, this is equivalent to
//  indices := d.cardIndices()
//  return indices[0] * 51! + indices[1] * 50! + ... + indices[51] * 0!
func (d Deck) Index() *big.Int {
	indices := d.cardIndices()
	retval := big.NewInt(0)
	for i, v := range indices {
		// retval *= i+1
		retval.Mul(retval, big.NewInt(int64(len(d)-i)))
		// retval += r
		retval.Add(retval, big.NewInt(int64(v)))
	}
	return retval
}

// fromCardIndices is the inverse of cardIndices.
func fromCardIndices(indices [DeckSize]Card) Deck {
	var d Deck
	var used [DeckSize]bool
	for i, c := range indices {
		for j, u := range used {
			if u {
				c++
			} else if int(c) == j {
				used[j] = true
				break
			}
		}

		d[i] = c
	}
	return d
}

// FromIndex computes the deck with the given index.
func FromIndex(index *big.Int) Deck {
	var indices [DeckSize]Card
	for i := range indices {
		// index, r = index / i+1, index % i+1
		r := big.NewInt(0)
		index, r = index.DivMod(index, big.NewInt(int64(i+1)), r)

		indices[len(indices)-i-1] = Card(r.Int64())
	}

	return fromCardIndices(indices)
}

func IndexEncode(d Deck) []byte {
	return d.Index().Bytes()
}

func IndexDecode(b []byte) Deck {
	return FromIndex(big.NewInt(0).SetBytes(b))
}
