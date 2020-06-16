package gates

// This gates intenrall uses an AND and NOT gate to implement the function of
// a NAND gate on an internal board
//
//    Pin1–––––|&&
//             |&& –––––Pin1––|>o––––Out
//    Pin2–––––|&&
//
type nandGate struct {
	id    string
	and   Gate
	not   Gate
	board Board
}

func (g *nandGate) ID() string {
	return g.id
}

// NAND creates a new NAND gate, which internall uses an AND and a
// NOT gate connected together
func NAND(board Board) Gate {
	// Creating the internal circuit on an innder board
	internalBoard := DefaultBoard()
	and := AND(internalBoard)
	not := NOT(internalBoard)
	internalBoard.Connect(and.Out(), not.Pin1())

	// Creating and adding to the parent board
	gate := nandGate{
		and:   and,
		not:   not,
		board: internalBoard,
	}
	// we simply add the gate without runing it on the parent board
	board.AddComponent(&gate)
	return &gate
}

func (g *nandGate) Pin1() chan int {
	return g.and.Pin1()
}

func (g *nandGate) Pin2() chan int {
	return g.and.Pin2()
}

func (g *nandGate) Out() chan int {
	return g.not.Out()
}

// Run does nothing since the base gates connected together
// do all the work
func (g *nandGate) Run() {}

// Stop the gate from doing it's job
func (g *nandGate) Stop() {
	g.board.Stop()
}

func (g *nandGate) CoreGatesCount() int {
	return g.board.CoreGatesCount()
}
