package circuits

import (
	"testing"

	"github.com/danishm/andornot/util"

	"github.com/danishm/andornot/gates"
)

func TestNBitAdder(t *testing.T) {

	// Creating a 4-Bit adder (fba)
	board := gates.DefaultBoard()
	fba := DefaultNBitAdder(board, 4)

	cases := []struct {
		a    string
		b    string
		s    string
		cout int
	}{
		{a: "0000", b: "0001", s: "0001", cout: 0},
		{a: "0001", b: "0001", s: "0010", cout: 0},
		{a: "0111", b: "0001", s: "1000", cout: 0},
		{a: "1100", b: "0011", s: "1111", cout: 0},
		{a: "1111", b: "0001", s: "0000", cout: 1},
		{a: "1111", b: "0010", s: "0001", cout: 1},
		{a: "1111", b: "0011", s: "0010", cout: 1},
	}

	for _, c := range cases {
		a := util.BinaryToIntArray(c.a)
		b := util.BinaryToIntArray(c.b)

		// Setting the input pins
		for i := range c.a {
			fba.A(i) <- a[i]
			fba.B(i) <- b[i]
		}
		fba.CIn() <- 0

		// Reading output
		out := make([]int, 4)
		for i := range out {
			out[i] = <-fba.S(i)
		}
		cout := <-fba.COut()

		// Compare
		s := util.IntArrayToBinary(out)
		if s != c.s {
			t.Logf(" %s + %s = expected:%s(caryr:%d) got:%s(carry:%d)", c.a, c.b, c.s, c.cout, s, cout)
			t.Fail()
		}

	}

	gatesCount := board.CoreGatesCount()
	expected := 44
	if gatesCount != expected {
		t.Logf("Expecting gate count to be %d, got %d", expected, gatesCount)
		t.Fail()
	}

	board.Stop()
}
