package gates

// This gates intenrally uses a mixed gates implementation of
// a XOR gate on an internal board
//
//    Pin1–––––––––––––|NAND
//     |               |NAND ––––|
//     |        |––––––|NAND     |––––|AND
//     |   Pin2-|                     |AND –––––– Out
//     |        |––––––|OR       |––––|AND
//     |               |OR ––––––|
//     |–––––––––––––––|OR
//
// Reference: https://en.wikipedia.org/wiki/XOR_gate
type xorMixedGate struct {
	pin1  chan int
	pin2  chan int
	nand  Gate
	or    Gate
	and   Gate
	board Board
}

// XOR creates a new XOR gate using internal circuitry
func XOR(board Board) Gate {

	// Creating input pins. The inputs in this case will act
	// as inputs to multiple gates
	pin1 := make(chan int)
	pin2 := make(chan int)

	// Creating an internal board with all the gates required
	internalBoard := DefaultBoard()
	nand := NAND(internalBoard)
	or := OR(internalBoard)
	and := AND(internalBoard)

	// Making all the connecteions
	internalBoard.Connect(pin1, nand.Pin1(), or.Pin1())
	internalBoard.Connect(pin2, nand.Pin2(), or.Pin2())
	internalBoard.Connect(nand.Out(), and.Pin1())
	internalBoard.Connect(or.Out(), and.Pin2())

	// Creating and adding to the parent board
	gate := xorMixedGate{
		pin1:  pin1,
		pin2:  pin2,
		nand:  nand,
		or:    or,
		and:   and,
		board: internalBoard,
	}
	board.AddComponent(&gate)
	return &gate
}

func (g *xorMixedGate) Pin1() chan int {
	return g.pin1
}

func (g *xorMixedGate) Pin2() chan int {
	return g.pin2
}

func (g *xorMixedGate) Out() chan int {
	return g.and.Out()
}

// Run does nothing since the base gates connected together
// do all the work
func (g *xorMixedGate) Run() {}

// Stop the gate from doing it's job
func (g *xorMixedGate) Stop() {
	g.board.Stop()
}
