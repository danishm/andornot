package gates

import (
	"testing"
)

func TestNOT(t *testing.T) {
	board := DefaultBoard()
	gate := NOT(board)

	cases := []struct {
		p1       int
		expected int
	}{
		{p1: 1, expected: 0},
		{p1: 0, expected: 1},
	}

	for _, c := range cases {
		gate.Pin1() <- c.p1
		actual := <-gate.Out()
		if actual != c.expected {
			t.Logf("NOT %d -> %d But Expected %d", c.p1, actual, c.expected)
			t.Fail()
		}
	}

	board.Stop()
}
