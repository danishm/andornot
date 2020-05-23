package circuits

import (
	"testing"
)

// TestXOR tests the XOR gate
func TestXOR(t *testing.T) {
	adder := DefaultFullAdder()

	// This is basically the full truth table of the adder
	cases := []struct {
		a    int
		b    int
		cin  int
		s    int
		cout int
	}{
		{cin: 0, a: 0, b: 0, s: 0, cout: 0},
		{cin: 0, a: 0, b: 1, s: 1, cout: 0},
		{cin: 0, a: 1, b: 0, s: 1, cout: 0},
		{cin: 0, a: 1, b: 1, s: 0, cout: 1},
		{cin: 1, a: 0, b: 0, s: 1, cout: 0},
		{cin: 1, a: 0, b: 1, s: 0, cout: 1},
		{cin: 1, a: 1, b: 0, s: 0, cout: 1},
		{cin: 1, a: 1, b: 1, s: 1, cout: 1},
	}

	for _, c := range cases {
		adder.A() <- c.a
		adder.B() <- c.b
		adder.CIn() <- c.cin
		actualS := <-adder.S()
		actualCout := <-adder.COut()
		if actualS != c.s || actualCout != c.cout {
			t.Logf("CIn:%d A:%d B:%d -> Sum:%d Cout:%d  expected Sum:%d Cout:%d", c.cin, c.a, c.b, actualS, actualCout, c.s, c.cout)
			t.Fail()
		}
	}

	gatesCount := adder.CoreGatesCount()
	if gatesCount != 11 {
		t.Logf("Expected gate count to be %d got %d", 11, gatesCount)
		t.Fail()
	}

	adder.Stop()
}
