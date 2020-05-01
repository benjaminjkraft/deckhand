package cards

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestEncoders(t *testing.T) {
	iterations := 10000
	tests := []struct {
		encoder func(Deck) []byte
		decoder func([]byte) Deck
		maxLen  int
	}{
		{IndexEncode, IndexDecode, 29},
		{ByteEncode, ByteDecode, 52},
		{AZEncode, AZDecode, 52},
	}

	for _, test := range tests {
		test := test
		name := filepath.Base(runtime.FuncForPC(reflect.ValueOf(test.encoder).Pointer()).Name())
		t.Run(name, func(t *testing.T) {
			rand.Seed(0) // make test reproducible

			for i := 0; i < iterations; i++ {
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

func TestShuffle(t *testing.T) {
	rand.Seed(0) // make test reproducible

	ordered := NewDeck()

	shuffled := NewDeck()
	shuffled.Shuffle()

	if ordered == shuffled {
		t.Errorf("shuffle didn't shuffle %v", ordered)
	}
}

func TestString(t *testing.T) {
	expected := "[🂡 🂢 🂣 🂤 🂥 🂦 🂧 🂨 🂩 🂪 🂫 🂭 🂮 🂱 🂲 🂳 🂴 🂵 🂶 🂷 🂸 🂹 🂺 🂻 🂽 🂾 🃁 🃂 🃃 🃄 🃅 🃆 🃇 🃈 🃉 🃊 🃋 🃍 🃎 🃑 🃒 🃓 🃔 🃕 🃖 🃗 🃘 🃙 🃚 🃛 🃝 🃞]"
	actual := fmt.Sprintf("%v", NewDeck())
	if actual != expected {
		t.Errorf("stringer incorrect:\ngot  %v\nwant %v", actual, expected)
	}
}
