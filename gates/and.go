package gates

// ANDGate represents and AND Gate
type ANDGate struct {
	Pin1 chan int
	Pin2 chan int
	Out  chan int
}

// Compute calculates the value
func (g *ANDGate) Compute() {
	x, y := <-g.Pin1, <-g.Pin2
	if x >= 1 && y >= 1 {
		g.Out <- 1
	}
	g.Out <- 0
}

func (g *ANDGate) Stop() {
	close(g.Pin1)
	close(g.Pin2)
	close(g.Out)
}
