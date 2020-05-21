package gates

import (
	"testing"
)

// TestNAND tests the NAND gate
func TestNAND(t *testing.T) {
	board := DefaultBoard()
	gate := NAND(board)

	cases := []struct {
		p1       int
		p2       int
		expected int
	}{
		{p1: 1, p2: 1, expected: 0},
		{p1: 1, p2: 0, expected: 1},
		{p1: 0, p2: 1, expected: 1},
		{p1: 0, p2: 0, expected: 1},
	}

	for _, c := range cases {
		gate.Pin1() <- c.p1
		gate.Pin2() <- c.p2
		actual := <-gate.Out()
		if actual != c.expected {
			t.Logf("%d NAND %d -> %d But Expected %d", c.p1, c.p2, actual, c.expected)
			t.Fail()
		}
	}

	board.Stop()
}
