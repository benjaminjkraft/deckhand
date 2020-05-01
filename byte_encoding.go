package cards

// ByteEncode encodes the i'th card as byte(i).
func ByteEncode(d Deck) []byte {
	retval := make([]byte, DeckSize)
	for i, c := range d {
		retval[i] = byte(c)
	}
	return retval
}

func ByteDecode(b []byte) Deck {
	var d Deck
	for i, c := range b {
		d[i] = Card(c)
	}
	return d
}

// AZEncode encodes the cards as A-Za-z.
func AZEncode(d Deck) []byte {
	retval := make([]byte, DeckSize)
	for i, c := range d {
		retval[i] = byte('A') + byte(c) + (byte(c)/26)*6
	}
	return retval
}

func AZDecode(b []byte) Deck {
	var d Deck
	for i, c := range b {
		d[i] = Card(byte(c) - byte('A') - ((byte(c)-byte('A'))/32)*6)
	}
	return d
}
