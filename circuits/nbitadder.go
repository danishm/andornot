package circuits

import (
	"github.com/danishm/andornot/core"
	"github.com/danishm/andornot/gates"
)

// NBitAdder defines the internface needed for an N-Bit adder
type NBitAdder interface {
	core.Component

	// Inputs
	A(i int) chan int
	B(i int) chan int
	CIn() chan int

	// Outputs
	S(i int) chan int
	COut() chan int
}

type nBitAdder struct {
	adders []FullAdder

	board gates.Board
}

// DefaultNBitAdder constructs an n-bit array on a provided board with the
// required number of bits
func DefaultNBitAdder(board gates.Board, bits int) NBitAdder {

	// Initializing all the adders
	ib := gates.DefaultBoard()
	var adders = make([]FullAdder, bits)
	for i := range adders {
		adders[i] = DefaultFullAdder(ib)
	}

	// Connect the adders
	for i := 0; i < (bits - 1); i++ {
		ib.Connect(adders[i].COut(), adders[i+1].CIn())
	}

	nba := nBitAdder{
		adders: adders,
		board:  ib,
	}
	board.AddComponent(&nba)
	return &nba
}

func (nba *nBitAdder) A(i int) chan int {
	return nba.adders[i].A()
}

func (nba *nBitAdder) B(i int) chan int {
	return nba.adders[i].B()
}

func (nba *nBitAdder) CIn() chan int {
	return nba.adders[0].CIn()
}

func (nba *nBitAdder) S(i int) chan int {
	return nba.adders[i].S()
}

func (nba *nBitAdder) COut() chan int {
	return nba.adders[len(nba.adders)-1].COut()
}

func (nba *nBitAdder) Run() {}

func (nba *nBitAdder) Stop() {
	nba.board.Stop()
}

func (nba *nBitAdder) CoreGatesCount() int {
	return nba.board.CoreGatesCount()

}
