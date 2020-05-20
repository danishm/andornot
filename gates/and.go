package gates

// ANDGate represents and AND Gate
type andGate struct {
	pin1 chan int
	pin2 chan int
	out  chan int
}

// AND creates a new AND gate
func AND(board Board) Gate {
	gate := andGate{
		pin1: make(chan int),
		pin2: make(chan int),
		out:  make(chan int),
	}
	board.Run(&gate)
	board.AddGate(&gate)
	return &gate
}

func (g *andGate) Pin1() chan int {
	return g.pin1
}

func (g *andGate) Pin2() chan int {
	return g.pin2
}

func (g *andGate) Out() chan int {
	return g.out
}

// Compute calculates the value
func (g *andGate) Compute() {
	x, y := <-g.pin1, <-g.pin2
	if x >= 1 && y >= 1 {
		g.out <- 1
	} else {
		g.out <- 0
	}
}

// Stop the gate from doing it's job
func (g *andGate) Stop() {
	close(g.pin1)
	close(g.pin2)
	close(g.out)
}
