package cards

import "math/big"

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
