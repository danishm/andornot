package gates

import "github.com/danishm/andornot/identity"

type notGate struct {
	id   string
	pin1 chan int
	out  chan int
}

func (g *notGate) ID() string {
	return g.id
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
		id:   identity.Get("not"),
		pin1: make(chan int),
		out:  make(chan int),
	}
	board.RunComponent(&gate)
	board.AddComponent(&gate)
	return &gate
}

// Run calculates the value
func (g *notGate) Run() {
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

func (g *notGate) CoreGatesCount() int {
	return 1
}
