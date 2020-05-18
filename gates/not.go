package gates

type NOTGate struct {
	Pin1 chan int
	Out  chan int
}

// Compute calculates the value
func (g *NOTGate) Compute() {
	x := <-g.Pin1
	if x < 1 {
		g.Out <- 1
	} else {
		g.Out <- 0
	}
}

func (g *NOTGate) Stop() {
	close(g.Pin1)
	close(g.Out)
}
