package cards

import (
	"math/big"
)

// Compute the index of this permutation in the lexicographic ordering of all
// permutations.
func (d Deck) Index() *big.Int {
	retval := big.NewInt(0)
	var indexes [DeckSize]Card
	var used [DeckSize]bool
	for i, c := range d {
		// retval *= i+1
		retval.Mul(retval, big.NewInt(int64(len(d)-i)))

		r := 0
		for _, u := range used[:c] {
			if !u {
				r++
			}
		}
		used[c] = true
		indexes[i] = Card(r)

		// retval += r
		retval.Add(retval, big.NewInt(int64(r)))
	}
	return retval
}

func FromIndex(index *big.Int) Deck {
	var indexes [DeckSize]Card
	for i := range indexes {
		// index, r = index / i, index % i
		r := big.NewInt(0)
		index, r = index.DivMod(index, big.NewInt(int64(i+1)), r)

		indexes[len(indexes)-i-1] = Card(r.Int64())
	}

	var d Deck
	var used [DeckSize]bool
	for i := range d {
		c := indexes[i]
		for j, u := range used {
			if u {
				c += 1
			} else if int(c) == j {
				used[j] = true
				break
			}
		}

		d[i] = c
	}
	return d
}

func IndexEncode(d Deck) []byte {
	return d.Index().Bytes()
}

func IndexDecode(b []byte) Deck {
	return FromIndex(big.NewInt(0).SetBytes(b))
}
