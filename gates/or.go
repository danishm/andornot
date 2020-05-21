package gates

type orGate struct {
	pin1 chan int
	pin2 chan int
	out  chan int
}

// OR creates a new OR gate
func OR(board Board) Gate {
	gate := orGate{
		pin1: make(chan int),
		pin2: make(chan int),
		out:  make(chan int),
	}
	board.Run(&gate)
	board.AddGate(&gate)
	return &gate
}

func (g *orGate) Pin1() chan int {
	return g.pin1
}

func (g *orGate) Pin2() chan int {
	return g.pin2
}

func (g *orGate) Out() chan int {
	return g.out
}

// Compute calculates the value
func (g *orGate) Compute() {
	x, ok1 := <-g.pin1
	y, ok2 := <-g.pin2
	if ok1 && ok2 {
		if x >= 1 || y >= 1 {
			g.out <- 1
		} else {
			g.out <- 0
		}
	}
}

// Stop the gate from doing it's job
func (g *orGate) Stop() {
	close(g.pin1)
	close(g.pin2)
	close(g.out)
}
