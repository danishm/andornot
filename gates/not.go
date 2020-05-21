package gates

type notGate struct {
	pin1 chan int
	out  chan int
}

func (g *notGate) Pin1() chan int {
	return g.pin1
}

func (g *notGate) Pin2() chan int {
	return nil
}

func (g *notGate) Out() chan int {
	return g.out
}

// NOT returns a NOT gate
func NOT(board Board) Gate {
	gate := notGate{
		pin1: make(chan int),
		out:  make(chan int),
	}
	board.Run(&gate)
	board.AddGate(&gate)
	return &gate
}

// Compute calculates the value
func (g *notGate) Compute() {
	x, ok := <-g.pin1
	if ok {
		if x < 1 {
			g.out <- 1
		} else {
			g.out <- 0
		}
	}
}

func (g *notGate) Stop() {
	close(g.pin1)
	close(g.out)
}
