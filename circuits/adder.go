package circuits

import (
	"github.com/danishm/andornot/core"
	"github.com/danishm/andornot/gates"
	"github.com/danishm/andornot/identity"
)

// FullAdder represents an interface exposed by a full binary adder
type FullAdder interface {
	core.Component

	// Inputs
	A() chan int
	B() chan int
	CIn() chan int

	// Outputs
	S() chan int
	COut() chan int
}

type fullAdder struct {
	id string

	a   chan int
	b   chan int
	cin chan int

	xor1 gates.Gate
	and1 gates.Gate
	xor2 gates.Gate
	and2 gates.Gate
	or   gates.Gate

	board gates.Board
}

// DefaultFullAdder creates and returns a full binary adder
func DefaultFullAdder(board gates.Board) FullAdder {

	a := make(chan int)
	b := make(chan int)
	cin := make(chan int)

	// Creating everything on an internal board
	ib := gates.DefaultBoard()
	xor1 := gates.XOR(ib)
	xor2 := gates.XOR(ib)
	and1 := gates.AND(ib)
	and2 := gates.AND(ib)
	or := gates.OR(ib)

	// Making all the connections based on full adder implementation at
	// https://www.electronics-tutorials.ws/combination/comb_7.html
	ib.Connect(a, xor1.Pin1(), and1.Pin1())
	ib.Connect(b, xor1.Pin2(), and1.Pin2())
	ib.Connect(xor1.Out(), xor2.Pin1(), and2.Pin1())
	ib.Connect(cin, xor2.Pin2(), and2.Pin2())
	ib.Connect(and1.Out(), or.Pin1())
	ib.Connect(and2.Out(), or.Pin2())

	adder := fullAdder{
		id:    identity.Get("adder"),
		a:     a,
		b:     b,
		cin:   cin,
		xor1:  xor1,
		and1:  and1,
		xor2:  xor2,
		and2:  and2,
		or:    or,
		board: ib,
	}

	board.AddComponent(&adder)
	return &adder
}

func (fa *fullAdder) ID() string {
	return fa.id
}

func (fa *fullAdder) A() chan int {
	return fa.a
}

func (fa *fullAdder) B() chan int {
	return fa.b
}

func (fa *fullAdder) CIn() chan int {
	return fa.cin
}

func (fa *fullAdder) S() chan int {
	return fa.xor2.Out()
}

func (fa *fullAdder) COut() chan int {
	return fa.or.Out()
}

func (fa *fullAdder) Run() {}

func (fa *fullAdder) Stop() {
	fa.board.Stop()
}

func (fa *fullAdder) CoreGatesCount() int {
	return fa.board.CoreGatesCount()
}
