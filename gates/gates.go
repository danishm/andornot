package gates

import "sync"

// Gate defines the interface a gate should support
type Gate interface {
	Compute()
	Stop()
}

type board struct {
	running bool
	mu      sync.Mutex
	gates   []Gate
}

// Run gets the gate running
func (b *board) run(g Gate) {
	go func() {
		for b.running {
			g.Compute()
		}
	}()
}

func (b *board) Stop() {

}

// AND creates and AND gate
func (b *board) AND() ANDGate {
	gate := ANDGate{
		Pin1: make(chan int),
		Pin2: make(chan int),
		Out:  make(chan int),
	}
	b.run(&gate)
	return gate
}

// NOT creates and NOT gate
func (b *board) NOT() NOTGate {
	gate := NOTGate{
		Pin1: make(chan int),
		Out:  make(chan int),
	}
	b.run(&gate)
	return gate
}

// Board creates a new circuit board to work with
func Board() board {
	return board{
		running: true,
	}
}
