package cards

import (
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestEncoders(t *testing.T) {
	tests := []struct {
		encoder    func(Deck) []byte
		decoder    func([]byte) Deck
		maxLen     int
		iterations int
	}{
		{IndexEncode, IndexDecode, 29, 10000},
	}

	for _, test := range tests {
		test := test
		name := filepath.Base(runtime.FuncForPC(reflect.ValueOf(test.encoder).Pointer()).Name())
		t.Run(name, func(t *testing.T) {
			rand.Seed(0) // make test reproducible

			for i := 0; i < test.iterations; i++ {
				d := NewDeck()
				d.Shuffle()

				encoded := test.encoder(d)
				if len(encoded) > test.maxLen {
					t.Errorf("encoding was %v > %v bytes on %v\n",
						len(encoded), test.maxLen, d)
				}

				decoded := test.decoder(encoded)
				if decoded != d {
					t.Errorf("round trip failed to decode %v:\noriginal  %v\nroundtrip %v\n", encoded, d, decoded)
				}
			}
		})
	}

}
